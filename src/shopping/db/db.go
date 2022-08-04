package db

import (
  "shopping/models"
)
// type Item struct {
//   Price float64
// }

func LoadItem(id int) *models.Item {
  return &models.Item{
    Price: 9.001,
  }
}
