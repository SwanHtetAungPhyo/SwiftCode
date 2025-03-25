package utils

import (
	utils2 "github.com/SwanHtetAungPhyo/swifcode/pkg/utils"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var mockCSVData = `COUNTRY ISO2 CODE,SWIFT CODE,CODE TYPE,NAME,ADDRESS,TOWN NAME,COUNTRY NAME,TIME ZONE
AL,AAISALTRXXX,BIC11,UNITED BANK OF ALBANIA SH.A,"HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",TIRANA,ALBANIA,Europe/Tirane
BG,ABIEBGS1XXX,BIC11,ABV INVESTMENTS LTD,"TSAR ASEN 20  VARNA, VARNA, 9002",VARNA,BULGARIA,Europe/Sofia
BG,ADCRBGS1XXX,BIC11,ADAMANT CAPITAL PARTNERS AD,"JAMES BOURCHIER BLVD 76A HILL TOWER SOFIA, SOFIA, 1421",SOFIA,BULGARIA,Europe/Sofia
`

func failOnErrTesting(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err.Error())
	}
}
func createCSVTEST(t *testing.T, fileName string) string {
	tmpFile, err := os.CreateTemp("", "testfile_*.csv") // This generates a valid file pattern
	if err != nil {
		failOnErrTesting(t, err)
	}
	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {
			failOnErrTesting(t, err)
		}
	}(tmpFile)

	_, err = tmpFile.Write([]byte(mockCSVData))
	if err != nil {
		failOnErrTesting(t, err)
	}
	return tmpFile.Name()
}

func TestLoadCSV(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	tempFilePath := createCSVTEST(t, mockCSVData)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			failOnErrTesting(t, err)
		}
	}(tempFilePath)
	data, err := utils2.loadCSV(tempFilePath, logger)
	if err != nil {
		failOnErrTesting(t, err)
	}
	if len(data) != 3 {
		t.Fatalf("Expected 9 records, got %d", len(data))
	}
	if data[0].SwiftCode != "AAISALTRXXX" || !data[0].IsHeadquarter {
		t.Fatalf("First record should be a headquarter with SwiftCode 'AAISALTRXXX'")
	}
	if data[1].SwiftCode != "ABIEBGS1XXX" || !data[1].IsHeadquarter {
		t.Fatalf("Second record should be a headquarter with SwiftCode 'ABIEBGS1XXX'")
	}
	if data[2].SwiftCode != "ADCRBGS1XXX" || !data[2].IsHeadquarter {
		t.Fatalf("Eighth record should be a headquarter with SwiftCode 'AKBKMTMTXXX'")
	}

}
