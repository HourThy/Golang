GetProCate();

function GetProCate() {
    $.ajax({
        url: "/getProCate",
        type: "GET",
        contentType: "application/json",
        success: function (response) {
            var len = response.length;
            $("#cbPRODCATE").empty();
            for (var i = 0; i < len; i++) {
                var CODE_CATE = response[i]['CODE_CATE'];
                var CODE = response[i]['CODE'];
                var CODE_EXT = response[i]['CODE_EXT'];
                var SUBITEM = response[i]['SUBITEM'];
                var CODE_DSC = response[i]['CODE_DSC'];
                var EXT_1 = response[i]['EXT_1'];
                var EXT_2 = response[i]['EXT_2'];
                var EXT_3 = response[i]['EXT_3'];
                var EXT_4 = response[i]['EXT_4'];
                var EXT_5 = response[i]['EXT_5'];

                $("#cbPRODCATE").append("<option value='" + CODE + "-" + SUBITEM + "'>" +
                    CODE + " | " + CODE_DSC + " | " + SUBITEM + "</option>");
            }
            localStorage.setItem("strCODE", CODE);
            localStorage.setItem("strSUBITEM", SUBITEM);
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}

GetDateTime();

function GetDateTime() {
    var dateTime = new Date($.now());
    var hours = (dateTime.getHours() < 10 ? "0" : "") + dateTime.getHours();
    var minutes = (dateTime.getMinutes() < 10 ? "0" : "") + dateTime.getMinutes();
    var seconds = (dateTime.getSeconds() < 10 ? "0" : "") + dateTime.getSeconds();

    var years = dateTime.getFullYear();
    var months = (dateTime.getMonth() < 10 ? "0" : "") + (parseInt(dateTime.getMonth()) + 1);
    var days = (dateTime.getDate() < 10 ? "0" : "") + dateTime.getDate();

    var fullDate =  years + "-" + months + "-" + days;
    var indat = years + months + days;
    var intime = hours + ":" + minutes + ":" + seconds;
    localStorage.setItem("indat", indat);
    localStorage.setItem("intime", intime);
    localStorage.setItem("fullDate", fullDate);
}

function GetLineID() {
    $.ajax({
        url: "/getLineID",
        type: "GET",
        contentType: "application/json",
        success: function (response) {
            var len = response.length;
            $("#cbLineID").empty();
            for (var i = 0; i < len; i++) {
                var CODE_CATE = response[i]['CODE_CATE'];
                var CODE = response[i]['CODE'];
                var CODE_EXT = response[i]['CODE_EXT'];
                var SUBITEM = response[i]['SUBITEM'];
                var CODE_DSC = response[i]['CODE_DSC'];
                var EXT_1 = response[i]['EXT_1'];
                var EXT_2 = response[i]['EXT_2'];
                var EXT_3 = response[i]['EXT_3'];
                var EXT_4 = response[i]['EXT_4'];
                var EXT_5 = response[i]['EXT_5'];

                $("#cbLineID").append("<option value='" + SUBITEM + "'>" +
                    SUBITEM + "</option>");
            }
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}


function GetOwnerID() {
    $.ajax({
        url: "/getOwnerID",
        type: "GET",
        contentType: "application/json",
        success: function (response) {
            var len = response.length;
            $("#cbOwerID").empty();
            for (var i = 0; i < len; i++) {
                var CODE_CATE = response[i]['CODE_CATE'];
                var CODE = response[i]['CODE'];
                var CODE_EXT = response[i]['CODE_EXT'];
                var SUBITEM = response[i]['SUBITEM'];
                var CODE_DSC = response[i]['CODE_DSC'];
                var EXT_1 = response[i]['EXT_1'];
                var EXT_2 = response[i]['EXT_2'];
                var EXT_3 = response[i]['EXT_3'];
                var EXT_4 = response[i]['EXT_4'];
                var EXT_5 = response[i]['EXT_5'];

                $("#cbOwerID").append("<option value='" +
                    CODE + " | " + SUBITEM + "'>" +
                    CODE + " | " + SUBITEM + "</option>");
            }

            localStorage.setItem("strOwnerID", SUBITEM);
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}


    $('#cbPRODCATE').on('change', function (e) {
        var optionSelected = $("option:selected", this);
        var proCateSelected = this.value; // 'D', 'V', 'Z', 'RUBB'
        //alert(valueSelected);
        sessionStorage.setItem("PROD_CATE_Selected", proCateSelected);
        $.ajax({
            url: "/postProCate",
            type: "POST",
            data: {
                'subitem': proCateSelected,
                'indat': localStorage.getItem("indat"),
                'intime': localStorage.getItem("intime"),
                'usr': localStorage.getItem('username')
            },
            // contentType: "application/json",
            success: function (response) {
                var len = response.length;
                $("#cbPRODID").empty();
                for (var i = 0; i < len; i++) {
                    var PRODUCT_ID = response[i]['PRODUCT_ID'];
                    var PRODUCT_DSC = response[i]['PRODUCT_DSC'];
                    $("#cbPRODID").append("<option value='" + PRODUCT_ID + "'>" +
                        PRODUCT_ID + "</option>");
                }
                //localStorage.setItem("headerText", );
            },
            failure: function (response) {
                alert("Get data error!");
                $("#table_postgre").hide();
            }
        });
    });

    $('#cbPRODID').on('change', function (e) {
        $("#tbOrder").fadeIn();
        var optionSelected = $("option:selected", this);
        var proIDSelected = this.value; // 'D', 'V', 'Z', 'RUBB'
        //alert(valueSelected);
        prodID = $("#cbPRODID").val();
        prodCate = $("#cbPRODCATE").val();
        prodCate = prodCate.substring(0, 1);
        startDate = $("#startDate").val();
        endDate = $("#endDate").val();
        sessionStorage.setItem("PROD_ID_Selected", proIDSelected);
        $.ajax({
            url: "/postProID",
            type: "POST",
            data: {
                'prodID': prodID,
                'prodCate': prodCate,
                'startDate': startDate,
                'endDate': endDate
            },
            // contentType: "application/json",
            success: function (response) {
                $("#tbOrder").find("tr:gt(0)").remove();
                $(response).each(function (index, value) {
                    var record =
                        "<tr><td>" +
                        (index + 1) +
                        "</td><td>" +
                        (value.LOT_ID) +
                        "</td><td>" +
                        value.NX_OPE_ID +
                        "</td><td>" +
                        value.RESV_EQPT_ID +
                        "</td><td>" +
                        value.LOT_STAT +
                        "</td><td>" +
                        value.RESV_DATE +
                        "</td><td>" +
                        value.RESV_SHIFT_SEQ +
                        "</td><td>" +
                        value.RESV_COMMENT +
                        "</td><td>" +
                        value.CLAIM_DATE +
                        "</td><td>" +
                        value.CLAIM_TIME +
                        "</td><td>" +
                        value.CLAIM_USER +
                        "</td><td>" +
                        value.PLAN_OPT_WEIGHT +
                        "</td><td>" +
                        value.SHT_CNT +
                        "</td><tr>";
                    record = record.substring(0, record.length - 9);
                    $("#tbOrder").append(record);
                   
                });
                $("#tbInserted").fadeOut();
            },
            failure: function (response) {
                alert("Get data error!");
            }
        });
    });

    $('#cbOwerID').on('change', function (e) {
        sessionStorage.removeItem("maxStr");
        var maxLOT_ID = "";
        $.ajax({
            url: "/getMaxLOTID",
            type: "GET",
            contentType: "application/json",
            success: function (response) {
                var len = response.length;
                for (var i = 0; i < len; i++) {
                    maxLOT_ID = response[i]['MAX(LOT_ID)'];
                }
                console.log("changed", maxLOT_ID);
                maxLOT_ID = maxLOT_ID.substring(9, 13);
                var maxLOT_ID_Count = parseInt(maxLOT_ID);
                if (maxLOT_ID_Count == null) {
                    maxLOT_ID_Count += 1;
                }
                var maxStr = "";
                if (maxLOT_ID_Count >= 1000) {
                    maxStr = maxLOT_ID_Count;
                } else if (maxLOT_ID_Count >= 100) {
                    maxStr = "0" + (maxLOT_ID_Count + 1);
                } else if (maxLOT_ID_Count >= 10) {
                    maxStr = "00" + (maxLOT_ID_Count + 1);
                } else if (maxLOT_ID_Count >= 0) {
                    maxStr = "000" + (maxLOT_ID_Count + 1);
                }
                sessionStorage.setItem("maxStr", maxStr);
            },
            failure: function (response) {
                alert("Get data error!");
            }
        });
    });

    $('#cbLineID').on('change', function (e) {
        var optionSelected = $("option:selected", this);
        var lineIDSelected = this.value; // 'D', 'V', 'Z', 'RUBB'
        var prodID = $("#cbPRODID").val(); //V-S579487PB150A
        //alert(valueSelected);
        //sessionStorage.setItem("PROD_ID_Selected", lineIDSelected);
        $.ajax({
            url: "/postLineID",
            type: "POST",
            data: {
                'lnid': lineIDSelected,
                'indat': localStorage.getItem("indat"),
                'intime': localStorage.getItem("intime"),
                'usr': localStorage.getItem('username'),
                'prodID': prodID
            },
            // contentType: "application/json",
            success: function (response) {
                var len = response.length;
                $("#cbLineID").empty();
                for (var i = 0; i < len; i++) {
                    // var PRODUCT_ID = response[i]['PRODUCT_ID'];
                    // var PRODUCT_DSC = response[i]['PRODUCT_DSC'];
                    var CODE = response[i]['CODE'];
                    var SUBITEM = response[i]['SUBITEM'];
                    var CODE_DSC = response[i]['CODE_DSC'];
                    $("#cbLineID").append("<option value='" + CODE + CODE_DSC + "'>" +
                        CODE + " | " + SUBITEM + " | " + CODE_DSC + " | " + "</option>");
                }
            },
            failure: function (response) {
                alert("Get data error!");
            }
        });
    });

CheckSession();
$("#uName").text(localStorage.getItem("username"));

function CheckSession() {
    if (localStorage.getItem("username") == null) {
        window.location.assign("/");
    }
}

function Add() {
    sessionStorage.removeItem("maxStr");
    GetLineID();
    GetOwnerID();
    GetDateTime();
    var proCate = sessionStorage.getItem('PROD_CATE_Selected');
    proCate = proCate.substring(0, 1);
    $("#txtProdCate").val(proCate)

    //var proID = sessionStorage.getItem('PROD_ID_Selected');
    //proID = proID.substring(0, 15);
    
    var proID = $("#cbPRODID").val();
    $("#txtProdID").val(proID)

    // Get Max work in day
    var maxLOT_ID = "";
    $.ajax({
        url: "/getMaxLOTID",
        type: "GET",
        contentType: "application/json",
        success: function (response) {
            var len = response.length;
            for (var i = 0; i < len; i++) {
                maxLOT_ID = response[i]['MAX(LOT_ID)'];
            }
            console.log(maxLOT_ID);
            maxLOT_ID = maxLOT_ID.substring(9, 13);
            var maxLOT_ID_Count = parseInt(maxLOT_ID);
            if (maxLOT_ID_Count == null) {
                maxLOT_ID_Count += 1;
            }
            var maxStr = "";
            if (maxLOT_ID_Count >= 1000) {
                maxStr = maxLOT_ID_Count;
            } else if (maxLOT_ID_Count >= 100) {
                maxStr = "0" + (maxLOT_ID_Count + 1);
            } else if (maxLOT_ID_Count >= 10) {
                maxStr = "00" + (maxLOT_ID_Count + 1);
            } else if (maxLOT_ID_Count >= 0) {
                maxStr = "000" + (maxLOT_ID_Count + 1);
            }
            sessionStorage.setItem("maxStr", maxStr);
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}

function onSubmit() {
    var count = 1;
    localStorage.setItem("count", count);
    var date = localStorage.getItem("indat");
    date = date.substring(2, 10);
    var ownerID = localStorage.getItem("strOwnerID");
    var saleOrder = "";
    var workOrderType = "NORM";
    var prodCate = $("#txtProdCate").val();
    var prodID = $("#txtProdID").val();
    var outPutWeight = $("#txtOutputWeight").val();
    var lineID = $("#cbLineID").val();
    var ownerID = $("#cbOwerID").val();
    ownerID = ownerID.split("|");
    for (var i = 0; i < ownerID.length; i++) {
        var strOwnerID = ownerID[1];
        strOwnerID = strOwnerID.trim();

        var getOwner = ownerID[0];
    }

    var claimUser = localStorage.getItem("username");
    var claimDate = localStorage.getItem("fullDate");
    var claimTime = localStorage.getItem("intime");

    //ownerid = ownerid.substring(0, 3);
    //alert(prodCate + " | " + prodID + " | " + outPutWeight + " | " + lineID + " | " + ownerID);
    //alertify.notify('Work added!', 'success', 1, function () {
        var maxStr = sessionStorage.getItem("maxStr");
        if(maxStr == ""){
            maxStr += "000" + 1;
        }
        var workOrderID = sessionStorage.getItem("PROD_CATE_Selected");
        workOrderID = workOrderID.replace("-", "");
        var strFormat = workOrderID + strOwnerID + date + maxStr;
        var workOrder = strFormat;
        var wordOrderDSC = strFormat;
        $("#tbOrder tbody").fadeOut(200);
        $('#tbInserted').fadeIn();
        $('#modal').modal('toggle');
        var markup = "<tr><td>" + saleOrder + "</td><td>" + workOrder + "</td><td>" + wordOrderDSC +
            "</td><td>" + workOrderType + "</td><td>" + prodCate + "</td><td>" + prodID + "</td><td>" + lineID + "</td><td>" + getOwner + "</td><td>" + claimUser + "</td><td>" +
            claimDate + "</td><td>" + claimTime + "</td></tr>";
        $("#tbInserted").append(markup);
        $("#tbOrder").hide(markup);
    //});
    var planOPTWEIGHT = $("#txtOutputWeight").val();
    var resvDATE = localStorage.getItem("fullDate");
    $.ajax({
        url: "/insertAEQPTRESV",
        type: "POST",
        data: {
            "LOT_ID": workOrder,
            "NX_OPE_ID": prodID,
            "CLAIM_USER": claimUser,
            "CLAIM_DATE": claimDate,
            "CLAIM_TIME": claimTime,
            "PLAN_OPT_WEIGHT": planOPTWEIGHT,
            "RESV_DATE": resvDATE
        },
        // contentType: "application/json",
        success: function (response) {
            //alert("Success");
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}