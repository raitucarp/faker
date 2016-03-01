package finance

type currency struct {
	Code   string
	Name   string
	Symbol string
}

var Currency = []currency{
	currency{
		Code:   "AED",
		Symbol: "",
		Name:   "UAE Dirham",
	},
	currency{
		Code:   "AFN",
		Symbol: "؋",
		Name:   "Afghani",
	},
	currency{
		Code:   "ALL",
		Symbol: "Lek",
		Name:   "Lek",
	},
	currency{
		Code:   "AMD",
		Symbol: "",
		Name:   "Armenian Dram",
	},
	currency{
		Code:   "ANG",
		Symbol: "ƒ",
		Name:   "Netherlands Antillian Guilder",
	},
	currency{
		Code:   "AOA",
		Symbol: "",
		Name:   "Kwanza",
	},
	currency{
		Code:   "ARS",
		Symbol: "$",
		Name:   "Argentine Peso",
	},
	currency{
		Code:   "AUD",
		Symbol: "$",
		Name:   "Australian Dollar",
	},
	currency{
		Code:   "AWG",
		Symbol: "ƒ",
		Name:   "Aruban Guilder",
	},
	currency{
		Code:   "AZN",
		Symbol: "ман",
		Name:   "Azerbaijanian Manat",
	},
	currency{
		Code:   "BAM",
		Symbol: "KM",
		Name:   "Convertible Marks",
	},
	currency{
		Code:   "BBD",
		Symbol: "$",
		Name:   "Barbados Dollar",
	},
	currency{
		Code:   "BDT",
		Symbol: "",
		Name:   "Taka",
	},
	currency{
		Code:   "BGN",
		Symbol: "лв",
		Name:   "Bulgarian Lev",
	},
	currency{
		Code:   "BHD",
		Symbol: "",
		Name:   "Bahraini Dinar",
	},
	currency{
		Code:   "BIF",
		Symbol: "",
		Name:   "Burundi Franc",
	},
	currency{
		Code:   "BMD",
		Symbol: "$",
		Name:   "Bermudian Dollar (customarily known as Bermuda Dollar)",
	},
	currency{
		Code:   "BND",
		Symbol: "$",
		Name:   "Brunei Dollar",
	},
	currency{
		Code:   "BOB BOV",
		Symbol: "$b",
		Name:   "Boliviano Mvdol",
	},
	currency{
		Code:   "BRL",
		Symbol: "R$",
		Name:   "Brazilian Real",
	},
	currency{
		Code:   "BSD",
		Symbol: "$",
		Name:   "Bahamian Dollar",
	},
	currency{
		Code:   "BWP",
		Symbol: "P",
		Name:   "Pula",
	},
	currency{
		Code:   "BYR",
		Symbol: "p.",
		Name:   "Belarussian Ruble",
	},
	currency{
		Code:   "BZD",
		Symbol: "BZ$",
		Name:   "Belize Dollar",
	},
	currency{
		Code:   "CAD",
		Symbol: "$",
		Name:   "Canadian Dollar",
	},
	currency{
		Code:   "CDF",
		Symbol: "",
		Name:   "Congolese Franc",
	},
	currency{
		Code:   "CHF",
		Symbol: "CHF",
		Name:   "Swiss Franc",
	},
	currency{
		Code:   "CLP CLF",
		Symbol: "$",
		Name:   "Chilean Peso Unidades de fomento",
	},
	currency{
		Code:   "CNY",
		Symbol: "¥",
		Name:   "Yuan Renminbi",
	},
	currency{
		Code:   "COP COU",
		Symbol: "$",
		Name:   "Colombian Peso Unidad de Valor Real",
	},
	currency{
		Code:   "CRC",
		Symbol: "₡",
		Name:   "Costa Rican Colon",
	},
	currency{
		Code:   "CUP CUC",
		Symbol: "₱",
		Name:   "Cuban Peso Peso Convertible",
	},
	currency{
		Code:   "CVE",
		Symbol: "",
		Name:   "Cape Verde Escudo",
	},
	currency{
		Code:   "CZK",
		Symbol: "Kč",
		Name:   "Czech Koruna",
	},
	currency{
		Code:   "DJF",
		Symbol: "",
		Name:   "Djibouti Franc",
	},
	currency{
		Code:   "DKK",
		Symbol: "kr",
		Name:   "Danish Krone",
	},
	currency{
		Code:   "DOP",
		Symbol: "RD$",
		Name:   "Dominican Peso",
	},
	currency{
		Code:   "DZD",
		Symbol: "",
		Name:   "Algerian Dinar",
	},
	currency{
		Code:   "EEK",
		Symbol: "",
		Name:   "Kroon",
	},
	currency{
		Code:   "EGP",
		Symbol: "£",
		Name:   "Egyptian Pound",
	},
	currency{
		Code:   "ERN",
		Symbol: "",
		Name:   "Nakfa",
	},
	currency{
		Code:   "ETB",
		Symbol: "",
		Name:   "Ethiopian Birr",
	},
	currency{
		Code:   "EUR",
		Symbol: "€",
		Name:   "Euro",
	},
	currency{
		Code:   "FJD",
		Symbol: "$",
		Name:   "Fiji Dollar",
	},
	currency{
		Code:   "FKP",
		Symbol: "£",
		Name:   "Falkland Islands Pound",
	},
	currency{
		Code:   "GBP",
		Symbol: "£",
		Name:   "Pound Sterling",
	},
	currency{
		Code:   "GEL",
		Symbol: "",
		Name:   "Lari",
	},
	currency{
		Code:   "GHS",
		Symbol: "",
		Name:   "Cedi",
	},
	currency{
		Code:   "GIP",
		Symbol: "£",
		Name:   "Gibraltar Pound",
	},
	currency{
		Code:   "GMD",
		Symbol: "",
		Name:   "Dalasi",
	},
	currency{
		Code:   "GNF",
		Symbol: "",
		Name:   "Guinea Franc",
	},
	currency{
		Code:   "GTQ",
		Symbol: "Q",
		Name:   "Quetzal",
	},
	currency{
		Code:   "GYD",
		Symbol: "$",
		Name:   "Guyana Dollar",
	},
	currency{
		Code:   "HKD",
		Symbol: "$",
		Name:   "Hong Kong Dollar",
	},
	currency{
		Code:   "HNL",
		Symbol: "L",
		Name:   "Lempira",
	},
	currency{
		Code:   "HRK",
		Symbol: "kn",
		Name:   "Croatian Kuna",
	},
	currency{
		Code:   "HTG USD",
		Symbol: "",
		Name:   "Gourde US Dollar",
	},
	currency{
		Code:   "HUF",
		Symbol: "Ft",
		Name:   "Forint",
	},
	currency{
		Code:   "IDR",
		Symbol: "Rp",
		Name:   "Rupiah",
	},
	currency{
		Code:   "ILS",
		Symbol: "₪",
		Name:   "New Israeli Sheqel",
	},
	currency{
		Code:   "INR",
		Symbol: "",
		Name:   "Indian Rupee",
	},
	currency{
		Code:   "INR BTN",
		Symbol: "",
		Name:   "Indian Rupee Ngultrum",
	},
	currency{
		Code:   "IQD",
		Symbol: "",
		Name:   "Iraqi Dinar",
	},
	currency{
		Code:   "IRR",
		Symbol: "﷼",
		Name:   "Iranian Rial",
	},
	currency{
		Code:   "ISK",
		Symbol: "kr",
		Name:   "Iceland Krona",
	},
	currency{
		Code:   "JMD",
		Symbol: "J$",
		Name:   "Jamaican Dollar",
	},
	currency{
		Code:   "JOD",
		Symbol: "",
		Name:   "Jordanian Dinar",
	},
	currency{
		Code:   "JPY",
		Symbol: "¥",
		Name:   "Yen",
	},
	currency{
		Code:   "KES",
		Symbol: "",
		Name:   "Kenyan Shilling",
	},
	currency{
		Code:   "KGS",
		Symbol: "лв",
		Name:   "Som",
	},
	currency{
		Code:   "KHR",
		Symbol: "៛",
		Name:   "Riel",
	},
	currency{
		Code:   "KMF",
		Symbol: "",
		Name:   "Comoro Franc",
	},
	currency{
		Code:   "KPW",
		Symbol: "₩",
		Name:   "North Korean Won",
	},
	currency{
		Code:   "KRW",
		Symbol: "₩",
		Name:   "Won",
	},
	currency{
		Code:   "KWD",
		Symbol: "",
		Name:   "Kuwaiti Dinar",
	},
	currency{
		Code:   "KYD",
		Symbol: "$",
		Name:   "Cayman Islands Dollar",
	},
	currency{
		Code:   "KZT",
		Symbol: "лв",
		Name:   "Tenge",
	},
	currency{
		Code:   "LAK",
		Symbol: "₭",
		Name:   "Kip",
	},
	currency{
		Code:   "LBP",
		Symbol: "£",
		Name:   "Lebanese Pound",
	},
	currency{
		Code:   "LKR",
		Symbol: "₨",
		Name:   "Sri Lanka Rupee",
	},
	currency{
		Code:   "LRD",
		Symbol: "$",
		Name:   "Liberian Dollar",
	},
	currency{
		Code:   "LTL",
		Symbol: "Lt",
		Name:   "Lithuanian Litas",
	},
	currency{
		Code:   "LVL",
		Symbol: "Ls",
		Name:   "Latvian Lats",
	},
	currency{
		Code:   "LYD",
		Symbol: "",
		Name:   "Libyan Dinar",
	},
	currency{
		Code:   "MAD",
		Symbol: "",
		Name:   "Moroccan Dirham",
	},
	currency{
		Code:   "MDL",
		Symbol: "",
		Name:   "Moldovan Leu",
	},
	currency{
		Code:   "MGA",
		Symbol: "",
		Name:   "Malagasy Ariary",
	},
	currency{
		Code:   "MKD",
		Symbol: "ден",
		Name:   "Denar",
	},
	currency{
		Code:   "MMK",
		Symbol: "",
		Name:   "Kyat",
	},
	currency{
		Code:   "MNT",
		Symbol: "₮",
		Name:   "Tugrik",
	},
	currency{
		Code:   "MOP",
		Symbol: "",
		Name:   "Pataca",
	},
	currency{
		Code:   "MRO",
		Symbol: "",
		Name:   "Ouguiya",
	},
	currency{
		Code:   "MUR",
		Symbol: "₨",
		Name:   "Mauritius Rupee",
	},
	currency{
		Code:   "MVR",
		Symbol: "",
		Name:   "Rufiyaa",
	},
	currency{
		Code:   "MWK",
		Symbol: "",
		Name:   "Kwacha",
	},
	currency{
		Code:   "MXN MXV",
		Symbol: "$",
		Name:   "Mexican Peso Mexican Unidad de Inversion (UDI)",
	},
	currency{
		Code:   "MYR",
		Symbol: "RM",
		Name:   "Malaysian Ringgit",
	},
	currency{
		Code:   "MZN",
		Symbol: "MT",
		Name:   "Metical",
	},
	currency{
		Code:   "NGN",
		Symbol: "₦",
		Name:   "Naira",
	},
	currency{
		Code:   "NIO",
		Symbol: "C$",
		Name:   "Cordoba Oro",
	},
	currency{
		Code:   "NOK",
		Symbol: "kr",
		Name:   "Norwegian Krone",
	},
	currency{
		Code:   "NPR",
		Symbol: "₨",
		Name:   "Nepalese Rupee",
	},
	currency{
		Code:   "NZD",
		Symbol: "$",
		Name:   "New Zealand Dollar",
	},
	currency{
		Code:   "OMR",
		Symbol: "﷼",
		Name:   "Rial Omani",
	},
	currency{
		Code:   "PAB USD",
		Symbol: "B/.",
		Name:   "Balboa US Dollar",
	},
	currency{
		Code:   "PEN",
		Symbol: "S/.",
		Name:   "Nuevo Sol",
	},
	currency{
		Code:   "PGK",
		Symbol: "",
		Name:   "Kina",
	},
	currency{
		Code:   "PHP",
		Symbol: "Php",
		Name:   "Philippine Peso",
	},
	currency{
		Code:   "PKR",
		Symbol: "₨",
		Name:   "Pakistan Rupee",
	},
	currency{
		Code:   "PLN",
		Symbol: "zł",
		Name:   "Zloty",
	},
	currency{
		Code:   "PYG",
		Symbol: "Gs",
		Name:   "Guarani",
	},
	currency{
		Code:   "QAR",
		Symbol: "﷼",
		Name:   "Qatari Rial",
	},
	currency{
		Code:   "RON",
		Symbol: "lei",
		Name:   "New Leu",
	},
	currency{
		Code:   "RSD",
		Symbol: "Дин.",
		Name:   "Serbian Dinar",
	},
	currency{
		Code:   "RUB",
		Symbol: "руб",
		Name:   "Russian Ruble",
	},
	currency{
		Code:   "RWF",
		Symbol: "",
		Name:   "Rwanda Franc",
	},
	currency{
		Code:   "SAR",
		Symbol: "﷼",
		Name:   "Saudi Riyal",
	},
	currency{
		Code:   "SBD",
		Symbol: "$",
		Name:   "Solomon Islands Dollar",
	},
	currency{
		Code:   "SCR",
		Symbol: "₨",
		Name:   "Seychelles Rupee",
	},
	currency{
		Code:   "SDG",
		Symbol: "",
		Name:   "Sudanese Pound",
	},
	currency{
		Code:   "SEK",
		Symbol: "kr",
		Name:   "Swedish Krona",
	},
	currency{
		Code:   "SGD",
		Symbol: "$",
		Name:   "Singapore Dollar",
	},
	currency{
		Code:   "SHP",
		Symbol: "£",
		Name:   "Saint Helena Pound",
	},
	currency{
		Code:   "SLL",
		Symbol: "",
		Name:   "Leone",
	},
	currency{
		Code:   "SOS",
		Symbol: "S",
		Name:   "Somali Shilling",
	},
	currency{
		Code:   "SRD",
		Symbol: "$",
		Name:   "Surinam Dollar",
	},
	currency{
		Code:   "STD",
		Symbol: "",
		Name:   "Dobra",
	},
	currency{
		Code:   "SVC USD",
		Symbol: "$",
		Name:   "El Salvador Colon US Dollar",
	},
	currency{
		Code:   "SYP",
		Symbol: "£",
		Name:   "Syrian Pound",
	},
	currency{
		Code:   "SZL",
		Symbol: "",
		Name:   "Lilangeni",
	},
	currency{
		Code:   "THB",
		Symbol: "฿",
		Name:   "Baht",
	},
	currency{
		Code:   "TJS",
		Symbol: "",
		Name:   "Somoni",
	},
	currency{
		Code:   "TMT",
		Symbol: "",
		Name:   "Manat",
	},
	currency{
		Code:   "TND",
		Symbol: "",
		Name:   "Tunisian Dinar",
	},
	currency{
		Code:   "TOP",
		Symbol: "",
		Name:   "Pa'anga",
	},
	currency{
		Code:   "TRY",
		Symbol: "TL",
		Name:   "Turkish Lira",
	},
	currency{
		Code:   "TTD",
		Symbol: "TT$",
		Name:   "Trinidad and Tobago Dollar",
	},
	currency{
		Code:   "TWD",
		Symbol: "NT$",
		Name:   "New Taiwan Dollar",
	},
	currency{
		Code:   "TZS",
		Symbol: "",
		Name:   "Tanzanian Shilling",
	},
	currency{
		Code:   "UAH",
		Symbol: "₴",
		Name:   "Hryvnia",
	},
	currency{
		Code:   "UGX",
		Symbol: "",
		Name:   "Uganda Shilling",
	},
	currency{
		Code:   "USD",
		Symbol: "$",
		Name:   "US Dollar",
	},
	currency{
		Code:   "UYU UYI",
		Symbol: "$U",
		Name:   "Peso Uruguayo Uruguay Peso en Unidades Indexadas",
	},
	currency{
		Code:   "UZS",
		Symbol: "лв",
		Name:   "Uzbekistan Sum",
	},
	currency{
		Code:   "VEF",
		Symbol: "Bs",
		Name:   "Bolivar Fuerte",
	},
	currency{
		Code:   "VND",
		Symbol: "₫",
		Name:   "Dong",
	},
	currency{
		Code:   "VUV",
		Symbol: "",
		Name:   "Vatu",
	},
	currency{
		Code:   "WST",
		Symbol: "",
		Name:   "Tala",
	},
	currency{
		Code:   "XAF",
		Symbol: "",
		Name:   "CFA Franc BEAC",
	},
	currency{
		Code:   "XAG",
		Symbol: "",
		Name:   "Silver",
	},
	currency{
		Code:   "XAU",
		Symbol: "",
		Name:   "Gold",
	},
	currency{
		Code:   "XBA",
		Symbol: "",
		Name:   "Bond Markets Units European Composite Unit (EURCO)",
	},
	currency{
		Code:   "XBB",
		Symbol: "",
		Name:   "European Monetary Unit (E.M.U.-6)",
	},
	currency{
		Code:   "XBC",
		Symbol: "",
		Name:   "European Unit of Account 9(E.U.A.-9)",
	},
	currency{
		Code:   "XBD",
		Symbol: "",
		Name:   "European Unit of Account 17(E.U.A.-17)",
	},
	currency{
		Code:   "XCD",
		Symbol: "$",
		Name:   "East Caribbean Dollar",
	},
	currency{
		Code:   "XDR",
		Symbol: "",
		Name:   "SDR",
	},
	currency{
		Code:   "XFU",
		Symbol: "",
		Name:   "UIC-Franc",
	},
	currency{
		Code:   "XOF",
		Symbol: "",
		Name:   "CFA Franc BCEAO",
	},
	currency{
		Code:   "XPD",
		Symbol: "",
		Name:   "Palladium",
	},
	currency{
		Code:   "XPF",
		Symbol: "",
		Name:   "CFP Franc",
	},
	currency{
		Code:   "XPT",
		Symbol: "",
		Name:   "Platinum",
	},
	currency{
		Code:   "XTS",
		Symbol: "",
		Name:   "Codes specifically reserved for testing purposes",
	},
	currency{
		Code:   "YER",
		Symbol: "﷼",
		Name:   "Yemeni Rial",
	},
	currency{
		Code:   "ZAR",
		Symbol: "R",
		Name:   "Rand",
	},
	currency{
		Code:   "ZAR LSL",
		Symbol: "",
		Name:   "Rand Loti",
	},
	currency{
		Code:   "ZAR NAD",
		Symbol: "",
		Name:   "Rand Namibia Dollar",
	},
	currency{
		Code:   "ZMK",
		Symbol: "",
		Name:   "Zambian Kwacha",
	},
	currency{
		Code:   "ZWL",
		Symbol: "",
		Name:   "Zimbabwe Dollar",
	},
}
