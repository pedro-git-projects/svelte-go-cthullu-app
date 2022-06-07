package data

type Description struct {
	StrDescription string `json:"str_description"`
	AppDescription string `json:"app_description"`
	ConDescription string `json:"con_description"`
	IntDescription string `json:"int_description"`
	SizDescription string `json:"siz_descrpition"`
	PowDescription string `json:"pow_description"`
	EduDescription string `json:"edu_description"`
	DexDescription string `json:"dex_description"`
}

func (i *Investigator) SetDescription() {
	i.strenghtDescription()
	i.appearenceDescription()
	i.constitutionDescription()
	i.intelligenceDescription()
	i.sizeDescription()
	i.powerDescription()
	i.dexDescription()
	i.educationDescription()
}

func (i *Investigator) strenghtDescription() {
	switch i.Str >= 0 {
	case i.Str >= 0 && i.Str < 15:
		i.Description.StrDescription = "debilitado: incapaz de levantar-se ou levantar uma xícara de chá."
	case i.Str >= 15 && i.Str < 50:
		i.Description.StrDescription = "insignificante, fraco."
	case i.Str >= 50 && i.Str < 90:
		i.Description.StrDescription = "mediana, para um humano."
	case i.Str >= 90 && i.Str < 99:
		i.Description.StrDescription = "certamente uma das pessoas mais fortes que a maioria conheceu."
	case i.Str >= 99:
		i.Description.SizDescription = "de nível mundial. Representa o ápice da força humana."
	}
}

func (i *Investigator) appearenceDescription() {
	switch i.App >= 0 {
	case i.App >= 0 && i.App < 15:
		i.Description.AppDescription = "tão repugnante que outros sofrem com medo, revolta ou pena."
	case i.App >= 15 && i.App < 50:
		i.Description.AppDescription = "feio, possivelmente desfigurado por conta de um acidente ou nascença."
	case i.App >= 50 && i.App < 90:
		i.Description.AppDescription = "média."
	case i.App >= 90 && i.App < 99:
		i.Description.AppDescription = "naturalmente magnético, uma das pessoas mais bonitas que a maioria conhecerá."
	case i.App >= 99:
		i.Description.AppDescription = "o auge do glamour e cool, provavelmente uma supermodelo ou estrela de cinema. Limite humano."
	}
}

func (i *Investigator) constitutionDescription() {
	switch i.Con >= 0 {
	case i.Con == 0:
		i.Description.ConDescription = "morto."
	case i.Con >= 0 && i.Con < 15:
		i.Description.ConDescription = "doente, propenso a mazelas prolongadas e provavelmente incapaz de operar sem assistência."
	case i.Con >= 15 && i.Con < 50:
		i.Description.ConDescription = "saúde fraca, você está propenso a crises de saúde, assim como  sentir dor."
	case i.Con >= 50 && i.Con < 90:
		i.Description.ConDescription = "saúde mediana."
	case i.Con >= 90 && i.Con < 99:
		i.Description.ConDescription = "robusto e resistente, ignora resfriados."
	case i.Con >= 99:
		i.Description.ConDescription = "saúde de ferro, capaz de suportar muita dor."
	}
}

func (i *Investigator) intelligenceDescription() {
	switch i.Int >= 0 {
	case i.Int == 0:
		i.Description.IntDescription = "sem intelecto. Incapaz de compreender o mundo ao seu redor."
	case i.Int >= 1 && i.Int < 50:
		i.Description.IntDescription = "aprendizado lento,  capaz de realizar apenas a matemática mais básica, ou ler livros de nível iniciante."
	case i.Int >= 50 && i.Int < 90:
		i.Description.IntDescription = "intelecto humano médio."
	case i.Int >= 90 && i.Int < 99:
		i.Description.IntDescription = "perspicaz, provavelmente capaz de compreender vários idiomas ou teoremas."
	case i.Int >= 99:
		i.Description.IntDescription = "genial, comparável ao Tesla. Limite do intelecto humano."
	}
}

func (i *Investigator) sizeDescription() {
	switch i.Siz >= 0 {
	case i.Siz == 1:
		i.Description.SizDescription = "bebê (1 - 12lbs)."
	case i.Siz > 1 && i.Siz <= 15:
		i.Description.SizDescription = "criança ou alguém de estatura muito baixa, como um anão (33lbs/15kg)."
	case i.Siz > 15 && i.Siz <= 65:
		i.Description.SizDescription = "tamanho humano médio (peso e altura moderados) (170lbs/75kg)."
	case i.Siz > 65 && i.Siz <= 80:
		i.Description.SizDescription = "muito alto, forte, ou obeso (240lbs/150 kg)."
	case i.Siz > 80 && i.Siz <= 99:
		i.Description.SizDescription = "superdimensionado em algum aspecto (330lbs/150kg)."
	}
}

func (i *Investigator) powerDescription() {
	switch i.Pow >= 0 {
	case i.Pow == 0:
		i.Description.PowDescription = "mente debilitada, sem força de vontade ou potencial mágico."
	case i.Pow >= 1 && i.Pow <= 15:
		i.Description.PowDescription = "de vontade fraca, facilmente dominado por aqueles com maior intelecto ou força de vontade.."
	case i.Pow > 15 && i.Pow <= 90:
		i.Description.PowDescription = "forte determinação, um alto potencial para se conectar com o invisível e mágico."
	case i.Pow == 100:
		i.Description.PowDescription = "vontade de ferro, você tem uma forte conexão com o 'reino' espiritual ou o mundo invisível."
	}
}

func (i *Investigator) dexDescription() {
	switch i.Dex >= 0 {
	case i.Dex == 0:
		i.Description.DexDescription = "incapaz de se mover sem ajuda."
	case i.Dex > 0 && i.Dex <= 15:
		i.Description.DexDescription = "lento, desajeitado com habilidades motoras pobres para manipulação fina."
	case i.Dex > 15 && i.Dex <= 50:
		i.Description.DexDescription = "destreza humana média."
	case i.Dex > 50 && i.Dex <= 90:
		i.Description.DexDescription = "rápido, ágil e capaz de realizar proezas de manipulação fina."
	case i.Dex > 90 && i.Dex <= 99:
		i.Description.DexDescription = "atleta de classe mundial. Máximo humano."
	}
}

func (i *Investigator) educationDescription() {
	switch i.Edu >= 0 {
	case i.Edu == 0:
		i.Description.EduDescription = "recém nascido."
	case i.Edu >= 1 && i.Edu <= 15:
		i.Description.EduDescription = "completamente ignorante em todos os sentidos."
	case i.Edu > 15 && i.Edu <= 60:
		i.Description.EduDescription = "ensino médio completo."
	case i.Edu > 60 && i.Edu <= 70:
		i.Description.EduDescription = "ensino superior completo."
	case i.Edu > 70 && i.Edu <= 80:
		i.Description.EduDescription = "mestre."
	case i.Edu > 80 && i.Edu <= 90:
		i.Description.EduDescription = "doutor, professor."
	case i.Edu > 90 && i.Edu <= 96:
		i.Description.EduDescription = "autoridade mundial em algum campo de estudo."
	case i.Edu > 96:
		i.Description.EduDescription = "atingiu o ápice da educação humana."
	}
}
