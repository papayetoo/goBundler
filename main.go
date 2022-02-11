package main

import (
	"encoding/json"
	"fmt"
	xci "goBundler/xcimageset"
	"io/fs"
	"io/ioutil"
	"strings"
)

func main() {
	// if len(os.Args) < 2 {
	// 	panic("goBunlder needs a Xcode Project directory")
	// }
	// rootDir := strings.Trim(os.Args[1], " ")

	// if strings.Compare(rootDir, "") == 0 {
	// 	panic("Invalid Directory")
	// }
	// files, err := ioutil.ReadDir(rootDir)

	// if err != nil {
	// 	return
	// }

	// // 파일 목록 읽어 오기
	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	// if len(os.Args) < 2 {
	// 	panic("Invalid directory")
	// }

	// rootDir := strings.Trim(os.Args[1], " ")

	targetDir := "/Users/papayetoo/Storage/Swift Project/Bundler/R/Images.xcassets"
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		panic("Erro reaindg dir")
	}
	
	xcassets := []xci.XCImageSet{}
	for _, f := range files {
		fname := f.Name()
		if strings.HasSuffix(fname, ".imageset") {
			fmt.Println(fname, targetDir+"/"+fname)
			imgSet, err := xcimageset(targetDir+"/"+fname)
			if err != nil {
				continue
			}
			xcassets = append(xcassets, imgSet)
		}		
	}

	for _, xcasset := range xcassets {
		fmt.Println(staticAssets(xcasset))
	}


	// f, err := ioutil.ReadFile("/Users/papayetoo/Storage/Swift Project/Bundler/R/Images.xcassets/karina.imageset/Contents.json")

	// if err != nil {
	// 	panic("Error Occured when reading imageset")
	// }

	// var imageSet xci.XCImageSet

	// json.Unmarshal(f, &imageSet)

	// fmt.Printf("%v\n", imageSet.Images)

	// fmt.Printf("AssetName : %s\n", imageSet.AssetName())

	
	// if err != nil {
	// 	panic("Writing R.h error")
	// }
	// fmt.Println("Make R.h completed")

	
	// for _, image := range imageSet.Images {
	// 	name := strings.Split(image.Filename, ".")
	// 	if len(name) == 0 {continue}
	// 	rsrc += fmt.Sprintf("\nstatic var %v: UIImage { .load(name: \"%v\") }\n", name[0], name[0])
	// }

	// fmt.Println(rsrc)

}


func writeHeaderFile() error { 
	header := `//
	//R.h
	//
	// Created By goBundler on 2022/02/11
	import <Foundation/Foundation.h>

	FOUNDATION_EXPORT double RVersionNumber;
	FOUNDATION_EXPORT const unsigned char RVersionString[];
	`
	err := ioutil.WriteFile("./R/R.h", []byte(header), fs.ModeAppend)
	return err
}

func writeSwiftFile() error {
	rsrc := `//
	//	R.swift
	// 
	//	Created by goBundler on
	
	import Foudation
	import UIKit
	
	public class R {
		static let bundler = Bundle(for: R.self)
	}
	
	public extension R {
		enum Images {}
	}
	`
	err := ioutil.WriteFile("./R/R.h", []byte(rsrc), fs.ModeAppend)
	return err
}

func xcimageset(fsdir string) (xci.XCImageSet, error) {

	var imgSet xci.XCImageSet
	f, err := ioutil.ReadFile(fsdir + "/" + "Contents.json")
	if err != nil {
		panic("Error reading Contents.json")
	}
	json.Unmarshal(f, &imgSet)

	return imgSet, err
}

func staticAssets(imgSet xci.XCImageSet) string {
	return fmt.Sprintf("\nstatic var %v: UIImage { .load(name: \"%v\")}\n", imgSet.AssetName(), imgSet.AssetName())
}