// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.bluice.com.br

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bluiceoficial/bludialogbox"

	c "bluiceoficial/micheckhash/controls"
	"os"
	"path/filepath"
	"strings"

	"github.com/bluiceoficial/blusmartflow"
)

func showGerarHash(a fyne.App) {
	c.LoadTranslations()
	w := a.NewWindow(c.T("Generate Hash"))
	w.Resize(fyne.NewSize(500, 300))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	flow := blusmartflow.New()

	lblTipoHash := widget.NewLabel(c.T("Hash Type"))
	lblTipoHash.TextStyle = fyne.TextStyle{Bold: true}
	sOptions := []string{"MD5", "SHA1", "SHA256", "SHA512"}
	cboTipoHash := widget.NewSelect(sOptions, func(string) {})
	cboTipoHash.PlaceHolder = "MD5"

	flow.AddRow(
		container.NewVBox(lblTipoHash, cboTipoHash),
	)

	lblArquivo := widget.NewLabel(c.T("Select file"))
	lblArquivo.TextStyle = fyne.TextStyle{Bold: true}
	txtArquivo := widget.NewEntry()
	txtArquivo.Disable()

	btnArquivo := widget.NewButton("...", func() {
		bludialogbox.NewOpenFile(a, c.T("Open File"), []string{}, false, func(filenames []string) {
			for _, filename := range filenames {
				txtArquivo.SetText(filename)
			}
		})
	})

	ctnArquivo := container.NewVBox(widget.NewLabel(""), btnArquivo)

	flow.AddColumn(
		container.NewVBox(lblArquivo, txtArquivo),
		ctnArquivo,
	)

	flow.Resize(ctnArquivo, 50, 0)

	var btnGerar *widget.Button
	var txtInfo *widget.Entry
	var btnSave *widget.Button
	var sTipoHash = ""
	btnGerar = widget.NewButton(c.T("Generate Hash"), func() {
		go func() {
			fyne.Do(func() {
				txtInfo.SetText(c.T("Generating Hash... Please wait!"))
				btnGerar.Disable()
			})

			sFilename := txtArquivo.Text

			sTipoHash = cboTipoHash.Selected
			if sTipoHash == "" {
				sTipoHash = "md5"
			} else {
				sTipoHash = strings.ToLower(sTipoHash)
			}

			file, _ := os.Open(sFilename)
			defer file.Close()

			hashInBytes := c.GetHash(sTipoHash, file)
			fyne.Do(func() {
				txtInfo.SetText(hashInBytes)
				btnGerar.Enable()
			})
		}()
	})

	flow.AddRow(layout.NewSpacer())
	flow.AddRow(layout.NewSpacer())
	flow.AddRow(layout.NewSpacer())

	flow.AddColumn(
		layout.NewSpacer(),
		btnGerar,
		layout.NewSpacer(),
	)

	flow.Gap(btnGerar, 0, 29)

	txtInfo = widget.NewEntry()
	txtInfo.Disable()

	btnSave = widget.NewButton(c.T("Save"), func() {
		sFilename := filepath.Base(txtArquivo.Text)
		sConteudo := fmt.Appendf(nil, "%s %s", txtInfo.Text, sFilename)
		err := os.WriteFile(txtArquivo.Text+"."+sTipoHash, sConteudo, 0644)
		
		if err != nil {
			bludialogbox.NewAlert(a, "MiCheckHash", err.Error(), true, "Ok")
		} else {
			bludialogbox.NewAlert(a, "MiCheckHash", c.T("File created successfully!"), false, "Ok")
		}
	})

	flow.AddColumn(
		txtInfo, btnSave,
	)

	flow.Resize(btnSave, 57, 36)

	w.SetContent(flow.Container)
	w.Show()
}
