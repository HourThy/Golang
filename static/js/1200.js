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