package dto

import "service-campaign-startup/model/entity"

type UserRegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EmailCheckRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type CampaignUri struct {
	ID int `uri:"id" binding:"required"`
}

type CampaignRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             entity.User
}

type CampaignImageRequest struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       entity.User
}

type TransactionUri struct {
	ID   int `uri:"id" binding:"required"`
	User entity.User
}
