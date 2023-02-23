package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/signintech/gopdf"
)

func main() {
	fileserver := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fileserver)
	http.HandleFunc("/upload", handleUpload)
	fmt.Print("The server is running at 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// / upload the image
func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Multipart form
	err := r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		http.Error(w, fmt.Sprintf("get form err: %s", err), http.StatusBadRequest)
		return
	}
	files := r.MultipartForm.File["files"]

	var imageFiles []*multipart.FileHeader
	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		if extension == ".jpg" || extension == ".jpeg" || extension == ".png" {
			imageFiles = append(imageFiles, file)
		} else {
			http.Error(w, "The extension is not matched!", http.StatusBadRequest)
			return
		}
	}

	if len(imageFiles) == 0 {
		http.Error(w, "Please Upload Image file.", http.StatusBadRequest)
		return
	}

	pdf := convertImagesToPDF(imageFiles, w)

	filename := imageFiles[0].Filename
	if err := pdf.WritePdf(filename + ".pdf"); err != nil {
		http.Error(w, fmt.Sprintf("write file err: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filename+".pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filename+".pdf")
}

// convert jpg, jpeg or png into pdf file
func convertImagesToPDF(files []*multipart.FileHeader, w http.ResponseWriter) *gopdf.GoPdf {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
			http.Error(w, "The extension is not matched!", http.StatusBadRequest)
			return nil
		}

		src, err := file.Open()
		if err != nil {
			http.Error(w, fmt.Sprintf("open file err: %s", err), http.StatusBadRequest)
			return nil
		}
		defer src.Close()

		image, err := gopdf.ImageHolderByReader(src)
		if err != nil {
			http.Error(w, fmt.Sprintf("load image file err: %s", err), http.StatusBadRequest)
			return nil
		}

		pdf.AddPage()

		pdf.ImageByHolder(image, 10, 17, nil)
	}

	return &pdf
}
