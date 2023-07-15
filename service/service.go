package service

import (
	"log"
	"menuProvider/model"
	"menuProvider/repository"
)

type IService interface {
	GetMenu() ([]*model.MenuItem, error)
}

type Service struct {
	r repository.IRepository
}

func NewService(r repository.IRepository) IService {
	return &Service{r: r}
}

func (s *Service) GetMenu() ([]*model.MenuItem, error) {
	menuItems, err := s.r.Find()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// arrange menu
	menuMap := make(map[int]*model.MenuItem)
	var rootItems []*model.MenuItem

	for _, menuItem := range menuItems {
		if menuItem.ParentID == 0 {
			rootItems = append(rootItems, menuItem)
		} else {
			parent := menuMap[menuItem.ParentID]
			if parent != nil {
				parent.Children = append(parent.Children, menuItem)
			}
		}

		menuMap[menuItem.ID] = menuItem
	}

	return rootItems, nil
}
