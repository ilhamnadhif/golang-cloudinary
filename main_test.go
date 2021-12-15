package golang_cloudinary

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
	"time"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetCloud() *cloudinary.Cloudinary {
	err := godotenv.Load()
	PanicIfError(err)

	cloudName := os.Getenv("CLOUDINARY_NAME")
	cloudKey := os.Getenv("CLOUDINARY_KEY")
	cloudSecret := os.Getenv("CLOUDINARY_SECRET")

	cld, errorCloud := cloudinary.NewFromParams(cloudName, cloudKey, cloudSecret)
	PanicIfError(errorCloud)
	return cld
}

func TestUpload(t *testing.T) {
	ctx := context.Background()
	cld := GetCloud()

	cloudFolder := os.Getenv("CLOUDINARY_FOLDER")

	res, err := cld.Upload.Upload(ctx, "ivana.jpg", uploader.UploadParams{
		PublicID: cloudFolder + "/ivanafelia" + "-" + strconv.Itoa(int(time.Now().Unix())),
		//Folder:   cloudFolder,
	})
	PanicIfError(err)
	fmt.Println(res)
}

func TestGetImage(t *testing.T) {
	ctx := context.Background()
	cld := GetCloud()

	res , err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "weplant/ivanafelia"})
	PanicIfError(err)
	fmt.Println(res.SecureURL)
}

func TestDeleteImage(t *testing.T) {
	ctx := context.Background()
	cld := GetCloud()

	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     "weplant/ivanafelia-1639562408",
	})
	PanicIfError(err)
	fmt.Println(res)
}
