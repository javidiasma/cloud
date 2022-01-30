package DAL

import (
	"cloudPart1/config"
	"cloudPart1/models/DAO"
)

func ReturnByRank(rank int) []DAO.History {
	var history []DAO.History
	config.DB.Where(&DAO.History{Rank: rank}).Find(&history)
	return history
}

func CreateDatabase(history DAO.History) DAO.History {
	tempHistory := DAO.History{
		Rank:         history.Rank,
		Name:         history.Name,
		Platform:     history.Platform,
		Year:         history.Year,
		Genre:        history.Genre,
		Publisher:    history.Publisher,
		NA_Sales:     history.NA_Sales,
		EU_Sales:     history.EU_Sales,
		JP_Sales:     history.JP_Sales,
		Other_Sales:  history.Other_Sales,
		Global_Sales: history.Global_Sales,
	}
	config.DB.Create(&tempHistory)
	return tempHistory
}

func GetDatabase() ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Model(&DAO.History{}).Find(&histories)
	if err.Error != nil {
		return nil, err.Error
	}
	return histories, nil
}

func GetByRank(rank int) ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Where(&DAO.History{Rank: rank}).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

//

func GetByName(name string) ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Where("name LIKE ?", "%"+name+"%").Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func GetByPlatform(platform string, num int) ([]DAO.History, error) {
	var histories []DAO.History
	//Limit(num)
	err := config.DB.Where(&DAO.History{Platform: platform}).Limit(num).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func GetByYear(year int, num int) ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Where(&DAO.History{Year: year}).Limit(num).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func GetByGenre(genre string, num int) ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Where(&DAO.History{Genre: genre}).Limit(num).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func GetByBestSellers(year int, platform string) ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Limit(5).Where("year = ? AND platform = ?", year, platform).Order("global_sales desc").Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func GetEUMoreThanNA() ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Where("eu_sales > na_sales").Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}
