GetBayID();

function GetBayID() {
    var BAY_ID ="";
    var OPI_FLG = "";
    var BAY_DSC = "";
    var ADDT_INFO_1 = "";
    var ADDT_INFO_2 = "";
    var ADDT_INFO_3 = "";
    $.ajax({
        url: "/getBayID",
        type: "GET",
        contentType: "application/json",
        success: function (response) {
            var len = response.length;
            for (var i = 0; i < len; i++) {
                BAY_ID = response[i]['BAY_ID'];
                OPI_FLG = response[i]['OPI_FLG'];
                BAY_DSC = response[i]['BAY_DSC'];
                ADDT_INFO_1 = response[i]['ADDT_INFO_1'];
                ADDT_INFO_2 = response[i]['ADDT_INFO_2'];
                ADDT_INFO_3 = response[i]['ADDT_INFO_3'];
                $("#cbBAYID").append("<option value='" + BAY_ID + "'>" + BAY_ID + "</option>");
            }
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}


$('#cbBAYID').on('change', function (e) {
    PostBayID();
});

$('#cbEQPTID').on('change', function (e) {
    PostMachineID();
});

function PostBayID(){
    var bay_id = $("#cbBAYID").val();
    $.ajax({
        url: "/postBayID",
        type: "POST",
        data: {'BAY_ID': bay_id},
        success: function (response) {
            var len = response.length;
            $("#cbEQPTID").empty();
            for (var i = 0; i < len; i++) {
                EQPT_ID = response[i]['EQPT_ID'];
                // OPI_FLG = response[i]['OPI_FLG'];
                // BAY_DSC = response[i]['BAY_DSC'];
                // ADDT_INFO_1 = response[i]['ADDT_INFO_1'];
                // ADDT_INFO_2 = response[i]['ADDT_INFO_2'];
                // ADDT_INFO_3 = response[i]['ADDT_INFO_3'];
                $("#cbEQPTID").append("<option value='" + EQPT_ID + "'>" + EQPT_ID + "</option>");
            }
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}

function PostMachineID(){
    var bay_id = $("#cbEQPTID").val();

    var LOT_ID = "";
    var NX_OPE_ID = "";
    var NX_OPE_VER = "";
    var PLANT_OPT_WEIGHT = "";
    var NX_OPE_NO = "";
    var RESV_DATE = "";
    var RESV_SHIFT_SEQ = "";
    $.ajax({
        url: "/getMachineID",
        type: "POST",
        data: {'BAY_ID': bay_id},
        success: function (response) {
            $("#table2").find("tr:gt(0)").remove();
            $(response).each(function (index, value) {
                var record =
                    "<tr><td>" +
                    "<input type='checkbox'></input>" +
                    "</td><td>" +
                    (index + 1) +
                    "</td><td>" +
                    (value.NX_OPE_NO) +
                    "</td><td>" +
                    value.NX_OPE_ID +
                    "</td><td>" +
                    value.NX_OPE_VER +
                    "</td><td>" +
                    value.PLAN_OPT_WEIGHT +
                    "</td><td>" +
                    value.NX_OPE_NO +
                    "</td><td>" +
                    value.RESV_DATE +
                    "</td><td>" +
                    value.RESV_SHIFT_SEQ +
                    "</td><tr>";
                record = record.substring(0, record.length - 9);
                $("#table2").append(record);
            });
        },
        failure: function (response) {
            alert("Get data error!");
        }
    });
}



$(document).ready(function() {
    $("#btnMoveLeft").on("click",function(){
            $('input:checked').each(function() {
            $(this).closest('tr').remove();
        });
    });
});