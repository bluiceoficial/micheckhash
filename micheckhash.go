// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.bluice.com.br

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	c "bluiceoficial/micheckhash/controls"
	"net/url"
	"os"
	"strings"

	"github.com/bluiceoficial/bludialogbox"
	"github.com/bluiceoficial/blusmartflow"
)

const VERSION_APP string = "7.0.0"

func main() {
	c.LoadTranslations()

	sIcon := fyne.NewStaticResource("micheckhash.png", resourceMicheckhashPngData)

	a := app.NewWithID("br.com.mugomes.micheckhash")
	a.SetIcon(sIcon)
	w := a.NewWindow("MiCheckHash")
	w.Resize(fyne.NewSize(500, 386))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	a.Settings().SetTheme(&myDarkTheme{})

	mnuTools := fyne.NewMenu(c.T("Tools"),
		fyne.NewMenuItem(
			c.T("Generate Hash"), func() {
				showGerarHash(a)
			}),
	)

	mnuAbout := fyne.NewMenu(c.T("About"),
		fyne.NewMenuItem(c.T("Check Update"), func() {
			url, err := url.Parse("https://github.com/mugomes/micheckhash/releases")
			if err == nil {
				a.OpenURL(url)
			}
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("Support MiCheckHash"), func() {
			url, err := url.Parse("https://mugomes.github.io/apoie.html")
			if err == nil {
				a.OpenURL(url)
			}
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(c.T("About MiCheckHash"), func() {
			showAbout(a)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(mnuTools, mnuAbout))

	flow := blusmartflow.New()

	lblTipo := widget.NewLabel(c.T("Hash Type"))
	lblTipo.TextStyle = fyne.TextStyle{Bold: true}

	sOptions := []string{"MD5", "SHA1", "SHA256", "SHA512"}
	txtTipo := widget.NewSelect(sOptions, func(string) {})
	txtTipo.PlaceHolder = "MD5"

	flow.AddRow(container.NewVBox(
		lblTipo, txtTipo,
	))

	lblArquivo := widget.NewLabel(c.T("File"))
	lblArquivo.TextStyle = fyne.TextStyle{Bold: true}
	lblArquivo.Move(fyne.NewPos(0, txtTipo.Position().Y+37))
	txtArquivo := widget.NewEntry()
	txtArquivo.SetPlaceHolder(c.T("Select file"))
	txtArquivo.Disable()
	btnArquivo := widget.NewButton("...", func() {
		bludialogbox.NewOpenFile(a, c.T("Open File"), []string{}, false, func(filenames []string) {
			for _, filename := range filenames {
				txtArquivo.SetText(filename)
			}
		})
	})
	btnArquivo.Resize(fyne.NewSize(30, 38.4))
	btnArquivo.Move(fyne.NewPos(txtArquivo.Size().Width+10, txtArquivo.Position().Y))

	ctnArquivo := container.NewVBox(widget.NewLabel(""), btnArquivo)
	flow.AddColumn(
		container.NewVBox(lblArquivo, txtArquivo),
		ctnArquivo,
	)

	flow.Resize(ctnArquivo, 50, 0)

	lblHash := widget.NewLabel(c.T("Type/Paste the Hash"))
	lblHash.TextStyle = fyne.TextStyle{Bold: true}
	txtHash := widget.NewEntry()

	flow.AddRow(
		container.NewVBox(lblHash, txtHash),
	)

	var lblInfo *widget.Label
	var btnCheck *widget.Button

	btnCheck = widget.NewButton(c.T("Check Now"), func() {
		go func() {
			fyne.Do(func() {
				lblInfo.SetText(c.T("Verifying Hash... Please wait!"))
				btnCheck.Disable()
			})

			sFilename := txtArquivo.Text
			sHash := txtHash.Text
			sTipoHash := txtTipo.Selected
			if sTipoHash == "" {
				sTipoHash = "md5"
			} else {
				sTipoHash = strings.ToLower(sTipoHash)
			}

			file, _ := os.Open(sFilename)
			defer file.Close()

			fileHash := c.GetHash(sTipoHash, file)

			fyne.Do(func() {
				lblInfo.SetText("")
			})
			if fileHash == sHash {
				fyne.Do(func() {
					bludialogbox.NewAlert(a, "MiCheckHash", c.T("Success!"), false, "Ok")
				})
			} else {
				fyne.Do(func() {
					bludialogbox.NewAlert(a, "MiCheckHash", c.T("Different!"), true, "Ok")
				})
			}

			fyne.Do(func() {
				btnCheck.Enable()
			})
		}()
	})

	flow.AddRow(layout.NewSpacer())
	flow.AddRow(layout.NewSpacer())
	flow.AddRow(layout.NewSpacer())

	flow.AddColumn(
		layout.NewSpacer(),
		btnCheck,
		layout.NewSpacer(),
	)

	lblInfo = widget.NewLabel("")

	flow.AddRow(lblInfo)

	w.SetContent(flow.Container)
	w.ShowAndRun()
}
