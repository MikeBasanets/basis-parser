package db

import (
	"context"
	"time"
)

func UpsertPants(item Pants) error {
	_, err := connectionPool.Exec(context.Background(),
		`upsert into pants (pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory, lastUpdated,
			fitType, legOpeningCm)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		item.PageUrl,
		item.ImageUrl,
		item.Color,
		item.Pattern,
		item.Description,
		item.Brand,
		item.Price,
		item.Season,
		item.Subcategory,
		item.LastUpdated,
		item.FitType,
		item.LegOpeningCm)
	return err
}

func UpsertShirt(item Shirt) error {
	_, err := connectionPool.Exec(context.Background(),
		`upsert into shirts (pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory, lastUpdated,
			fitType, lengthCm, sleeveLengthCm, collarOrCutout)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		item.PageUrl,
		item.ImageUrl,
		item.Color,
		item.Pattern,
		item.Description,
		item.Brand,
		item.Price,
		item.Season,
		item.Subcategory,
		item.LastUpdated,
		item.FitType,
		item.LengthCm,
		item.SleeveLengthCm,
		item.CollarOrCutout)
	return err
}

func UpsertOuterwear(item Outerwear) error {
	_, err := connectionPool.Exec(context.Background(),
		`upsert into outerwear (pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory, lastUpdated,
			hoodType, lengthCm, sleeveLengthCm, insulationComposition)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		item.PageUrl,
		item.ImageUrl,
		item.Color,
		item.Pattern,
		item.Description,
		item.Brand,
		item.Price,
		item.Season,
		item.Subcategory,
		item.LastUpdated,
		item.HoodType,
		item.LengthCm,
		item.SleeveLengthCm,
		item.InsulationComposition)
	return err
}

func RemoveClothingUpdatedBefore(t time.Time) error {
	_, err := connectionPool.Exec(context.Background(), `delete from pants where lastUpdated < $1`, t)
	if err != nil {
		return err
	}
	_, err = connectionPool.Exec(context.Background(), `delete from shirts where lastUpdated < $1`, t)
	if err != nil {
		return err
	}
	_, err = connectionPool.Exec(context.Background(), `delete from outerwear where lastUpdated < $1`, t)
	return err
}
