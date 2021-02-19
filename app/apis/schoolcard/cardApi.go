package schoolcard

import "funnel/app/apis"

var CardHome = apis.CARD_URL + "default.aspx"
var CardBalance = apis.CARD_URL + "Cardholder/AccBalance.aspx"

var CardHistoryQuery = apis.CARD_URL + "Cardholder/Queryhistory.aspx"
var CardHistory = apis.CARD_URL + "Cardholder/QueryhistoryDetailFrame.aspx"
var CardToday = apis.CARD_URL + "Cardholder/QueryCurrDetailFrame.aspx"
