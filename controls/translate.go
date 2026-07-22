// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.bluice.com.br

package controls

import "github.com/bluiceoficial/blulang"

func LoadTranslations() {
	lang := blulang.GetLang()
	if lang == "pt" {
		blulang.Set("File", "Arquivo")
		blulang.Set("Save", "Salvar")
		blulang.Set("Hash Type", "Tipo de Hash")
		blulang.Set("Select file", "Selecione o arquivo")
		blulang.Set("Generate Hash", "Gerar Hash")
		blulang.Set("Success!", "Sucesso!")
		blulang.Set("Different!", "Diferente")
		blulang.Set("Tools", "Ferramentas")
		blulang.Set("About", "Sobre")
		blulang.Set("Check Update", "Verificar Atualização")
		blulang.Set("Support MiCheckHash", "Apoie MiCheckHash")
		blulang.Set("About MiCheckHash", "Sobre MiCheckHash")
		blulang.Set("Type/Paste the Hash", "Digite/Cole o Hash")
		blulang.Set("Check Now", "Verificar Agora")
		blulang.Set("Verifying Hash... Please wait!", "Verificando Hash... Aguarde!")
		blulang.Set("Generating Hash... Please wait!", "Gerando Hash... Aguarde!")
		blulang.Set("File created successfully!", "Arquivo criado com sucesso!")
	}
}

func T(key string, args ...interface{}) string {
	return blulang.T(key, args...)
}
