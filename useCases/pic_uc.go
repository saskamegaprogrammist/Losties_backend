package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type PicUC struct {
	AdsDB *database.AdsDB
	UsersDB *database.UsersDB
	FilenameDB *database.FilenameDB

}


func (picUC *PicUC) NewAdPic(ad *models.Ad, file multipart.File, id string) (bool, error) {
	err := picUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	var newAd models.Ad
	newAd.Id = ad.UserId
	err = picUC.AdsDB.GetAdById(&newAd)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID || newAd.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" && fileType != "image/jpg" &&
		fileType != "image/gif" && fileType != "image/png" {
		return false, err
	}
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		return false, err
	}

	newPath := filepath.Join("./pics/", "ad_pic_" + id +fileEndings[0])
	newFile, err := os.Create(newPath)
	if err != nil {
		return false, err
	}
	var filenameAd models.FilenameAd
	filenameAd.Ad = ad.Id
	filenameAd.Filename = newPath
	err = picUC.FilenameDB.InsertFilenameAd(&filenameAd)
	if err != nil {
		return false, err
	}

	defer newFile.Close()
	_, err = newFile.Write(fileBytes)
	if  err != nil {
		return false, err
	}
	return false, nil
}

func (picUC *PicUC) GetAdPic (ad *models.Ad) (bool, *os.File, error) {
	err := picUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, nil, err
	}
	if ad.Id == utils.ERROR_ID {
		return true, nil, fmt.Errorf("this ad doesn't exist")
	}
	filenameAd, err := picUC.FilenameDB.GetFilenameByAdId(ad.Id)
	var filenameString string
	if err != nil {
		filenameString = "./pics/cat.jpg"
	} else {
		filenameString = filenameAd.Filename
	}
	file, err := os.Open(filenameString)
	if err != nil {
		return false, nil, err
	}
	return false, file, nil
}

func (picUC *PicUC) NewUserPic(user *models.User, file multipart.File, id string) (bool, error) {
	err := picUC.UsersDB.GetUserById(user)
	if err != nil {
		return false, err
	}
	if user.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this user doesn't exist")
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" && fileType != "image/jpg" &&
		fileType != "image/gif" && fileType != "image/png" {
		return false, err
	}
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		return false, err
	}

	newPath := filepath.Join("./pics/", "user_pic_" + id +fileEndings[0])
	newFile, err := os.Create(newPath)
	if err != nil {
		return false, err
	}
	var filenameUser models.FilenameUser
	filenameUser.User = user.Id
	filenameUser.Filename = newPath
	err = picUC.FilenameDB.InsertFilenameUser(&filenameUser)
	if err != nil {
		return false, err
	}

	defer newFile.Close()
	_, err = newFile.Write(fileBytes)
	if  err != nil {
		return false, err
	}
	return false, nil
}

func (picUC *PicUC) GetUserPic (user *models.User) (bool, *os.File, error) {
	err := picUC.UsersDB.GetUserById(user)
	if err != nil {
		return false, nil, err
	}
	if user.Id == utils.ERROR_ID {
		return true, nil, fmt.Errorf("this user doesn't exist")
	}
	filenameUser, err := picUC.FilenameDB.GetFilenameByUserId(user.Id)
	var filenameString string
	if err != nil {
		filenameString = "./pics/user.png"
	} else {
		filenameString = filenameUser.Filename
	}
	log.Println(filenameString)
	file, err := os.Open(filenameString)
	if err != nil {
		return false, nil, err
	}
	return false, file, nil
}