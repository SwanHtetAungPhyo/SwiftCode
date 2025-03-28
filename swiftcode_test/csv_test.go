// Package swiftcode_test contains unit tests for CSV parsing functionality.
// It validates the LoadCSV function from the utils package.
package swiftcode_test

import (
	"os"
	"testing"

	"github.com/SwanHtetAungPhyo/swifcode/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test data constants
const (
	expectedRecords = 3
)

// mockCSVData contains sample CSV data for testing
var mockCSVData = `COUNTRY ISO2 CODE,SWIFT CODE,CODE TYPE,NAME,ADDRESS,TOWN NAME,COUNTRY NAME,TIME ZONE
AL,AAISALTRXXX,BIC11,UNITED BANK OF ALBANIA SH.A,"ADDRESS1",TIRANA,ALBANIA,Europe/Tirane
BG,ABIEBGS1XXX,BIC11,ABV INVESTMENTS LTD,"ADDRESS2",VARNA,BULGARIA,Europe/Sofia
BG,ADCRBGS1XXX,BIC11,ADAMANT CAPITAL PARTNERS AD,"ADDRESS3",SOFIA,BULGARIA,Europe/Sofia
`

// TestLoadCSV validates CSV parsing functionality
func TestLoadCSV(t *testing.T) {
	// Setup test logger
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	// Create temp file
	tmpFile := createTempCSV(t)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			logger.WithError(err).Errorln("Failed to remove temporary file")
		}
	}(tmpFile)

	// Execute test
	t.Run("should parse valid CSV correctly", func(t *testing.T) {
		records, err := utils.LoadCSV(tmpFile, logger)

		// Validate results
		require.NoError(t, err)
		assert.Len(t, records, expectedRecords)

		testCases := []struct {
			index         int
			swiftCode     string
			isHeadquarter bool
		}{
			{0, "AAISALTRXXX", true},
			{1, "ABIEBGS1XXX", true},
			{2, "ADCRBGS1XXX", true},
		}

		for _, tc := range testCases {
			assert.Equal(t, tc.swiftCode, records[tc.index].SwiftCode)
			assert.True(t, records[tc.index].IsHeadquarter)
		}
	})
}

// createTempCSV generates a temporary CSV file for testing
func createTempCSV(t *testing.T) string {
	tmpFile, err := os.CreateTemp("", "test-*.csv")
	require.NoError(t, err, "Should create temp file")
	_, err = tmpFile.WriteString(mockCSVData)
	require.NoError(t, err, "Should write test data")

	err = tmpFile.Close()
	require.NoError(t, err, "Should close temp file")

	return tmpFile.Name()
}
