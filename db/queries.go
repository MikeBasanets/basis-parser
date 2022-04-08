package db

import (
	"context"
)

func UpsertPants(item Pants) error {
	_, err := connectionPool.Exec(context.Background(),
		`upsert into pants (pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory,
			fitType, legOpeningCm)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		item.PageUrl,
		item.ImageUrl,
		item.Color,
		item.Pattern,
		item.Description,
		item.Brand,
		item.Price,
		item.Season,
		item.Subcategory,
		item.FitType,
		item.LegOpeningCm)
	return err
}

func UpsertShirt(item Shirt) error {
	_, err := connectionPool.Exec(context.Background(),
		`upsert into shirts (pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory,
			fitType, lengthCm, sleeveLengthCm, collarOrCutout)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		item.PageUrl,
		item.ImageUrl,
		item.Color,
		item.Pattern,
		item.Description,
		item.Brand,
		item.Price,
		item.Season,
		item.Subcategory,
		item.FitType,
		item.LengthCm,
		item.SleeveLengthCm,
		item.CollarOrCutout)
	return err
}

func UpsertOuterwear(item Outerwear) error {
	_, err := connectionPool.Exec(context.Background(),
		`upsert into outerwear (pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory,
			hoodType, lengthCm, sleeveLengthCm, insulationComposition)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		item.PageUrl,
		item.ImageUrl,
		item.Color,
		item.Pattern,
		item.Description,
		item.Brand,
		item.Price,
		item.Season,
		item.Subcategory,
		item.HoodType,
		item.LengthCm,
		item.SleeveLengthCm,
		item.InsulationComposition)
	return err
}
