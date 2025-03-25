package services

//// DataProcessing handles the entire pipeline
//func DataProcessing(log *logrus.Logger) {
//	filePath := "/Users/swanhtet1aungphyo/IdeaProjects/SwiftCode/data/swif_codes.csv"
//
//	// Parse & Extract Data
//	countries, bankDtos := ParseAndProcessData(filePath, log)
//
//	// Migrate & Insert Countries
//	if err := repo.MigrateTables(log); err != nil {
//		log.Fatal("Migration failed:", err)
//	}
//	repo.InsertCountries(countries, log)
//
//	// Fetch Country IDs
//	countryIDMap := repo.FetchCountryIDs(countries, log)
//	if len(countryIDMap) == 0 {
//		log.Fatal("No ISO codes found in DB")
//	}
//
//	// Map Banks to Country IDs & Insert
//	banks := ConvertToBankDetails(AssignCountryIDs(bankDtos, countryIDMap, log))
//	repo.InsertBankDetails(banks, log)
//
//	log.Info("Data processing complete.")
//}
