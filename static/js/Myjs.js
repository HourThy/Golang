CheckSession();

function CheckSession() {
    if (localStorage.getItem("username") == null) {
        window.location.assign("/");
    }
}

function UpdateRubmats() {
    var id = $("#updateModal #id").val();
    var machno = $("#updateModal #findMachno").val();
    var matetype = $("#updateModal #findMatetype").val();
    var matecode = $("#updateModal #findMatecode").val();
    var matename = $("#updateModal #findMatename").val();
    var intime = $("#updateModal #intime").val();
    var indat = $("#updateModal #indat").val();
    var usrno = $("#updateModal #usrno").val();
    var state = $("#updateModal #state").val();
    $.ajax({
        url: "/updateRubmat",
        type: "POST",
        data: {
            'id': id,
            'machno': machno,
            'matetype': matetype,
            'matecode': matecode,
            'matename': matename,
            'intime': intime,
            'indat': indat,
            'usrno': usrno,
            'state': state
        },
        success: function (response) {
            var notification = alertify.notify('Update successful!', 'success', 0.1, function () {
                //console.log("success");
                $("#updateModal").load();
            });
        },
        failure: function (response) {
            //$.alert(response.d);
            var notification = alertify.notify('Update failed', 'danger', 1, function () {
                console.log('dismissed');
            });
        }
    });
}

function GetRubmats() {
    $("#table_rubmats").fadeIn();
    $("#table_add").fadeOut(200);
    $("#btnOK").fadeOut(200);
    $("#insertTitle").fadeOut(200);
    $("#rows_search").fadeIn(200);
    $("#addedTittle").fadeOut(200);
    $("#table_added").fadeOut(200);
    var ID = $("#id").val();

    var Machno = $("#machno").val();
    var Matetype = $("#matetype").val();
    var Matecode = $("#matecode").val();
    var Matename = $("#matename").val();
    var Intime = $("#intime").val();
    var Indat = $("#indat").val();
    var Usrno = $("#usrno").val();
    var State = $("#state").val();
    var FindMachno = $("#findMachno").val();
    var FindMateType = $("#findMateType").val();
    var FindMateCode = $("#findMateCode").val();
    var FindMateName = $("#findMateName").val();
    $.ajax({
        url: "/getRubmats",
        type: "GET",
        data: {
            "id": ID,
            "machno": Machno,
            "matetype": Matetype,
            "matecode": Matecode,
            "matename": Matename,
            "intime": Intime,
            "indat": Indat,
            "usrno": Usrno,
            "state": State,
            "findMachno": FindMachno,
            "findMateType": FindMateType,
            "findMateCode": FindMateCode,
            "findMateName": FindMateName,
        },
        success: function (response) {
            // Xóa toàn bộ dòng trong table và thêm lại dòng mới

            $("#table_rubmats").find("tr:gt(0)").remove();
            $(response).each(function (index, value) {
                if ($("#table_rubmats tr").val() != null) {
                    var record =
                        "<tr><td>" +
                        (index + 1) +
                        "</td><td style='display: none' id='mongoID'>" +
                        (value.ID) +
                        "</td><td>" +
                        value.Machno +
                        "</td><td>" +
                        value.Matetype +
                        "</td><td>" +
                        value.Matecode +
                        "</td><td>" +
                        value.Matename +
                        "</td><td>" +
                        value.Intime +
                        "</td><td>" +
                        value.Indat +
                        "</td><td>" +
                        value.Usrno +
                        "</td><td>" +
                        value.State +
                        "</td><td>" +
                        "<button class='btn btn-warning' id='btnModifield' onclick='GetModified()' style='margin-right: 5px'>Modified</button>" +
                        "<button class='btn btn-danger' id='btnDelete' style='margin-right: 5px; width: 100px;' onclick='DeleteRubmats()'>Delete</button >" +
                        // "</td><td>" +

                        // "</td><td>" +
                        "</td><tr>";
                    record = record.substring(0, record.length - 9);
                    $("table").append(record);
                }
            });
        },
        error: function (response) {
            alert("Not found!\n找不到任何文件!");
            $("#table_rubmats").hide();
        }
    });
}

function InsertRubmats() {
    debugger;
    var machno = $("#findMachno").val();
    var matetype = $("#findMateType").val();
    var matecode = $("#findMateCode").val();
    var matename = $("#findMateName").val();
    var intime = localStorage.getItem("intime");
    var indat = localStorage.getItem("indat");
    var usrno = localStorage.getItem("username");
    var state = "0";
    if (machno == "" || matetype == "" || matecode == "" || matename == "") {
        alert("All field required!\n有字段必填！");
        return false;
    }
    $.ajax({
        url: "/insertRubmats",
        type: "POST",
        data: {
            'machno': machno,
            'matetype': matetype,
            'matecode': matecode,
            'matename': matename,
            'intime': intime,
            'indat': indat,
            'usrno': usrno,
            'state': state
        },
        success: function (response) {
            alertify.notify('User added successful', 'success', 1, function () {
                $("#table_added").fadeIn(200);
                $("#addedTittle").fadeIn(200);
                $("#findMachno").val(machno);
                $("#findMateType").val(matetype);
                $("#findMateCode").val(matecode);
                $("#findMateName").val(matename);
                $("#inserted_intime").val(intime);
                $("#inserted_indat").val(indat);
                $("#inserted_usrno").val(usrno);
                $("#inserted_state").val(state);
                $("#findMachno").val('');
                $("#findMateType").val('');
                $("#findMateName").val('');
                $("#findMateCode").val('');
                GetDateTime();
                $("#Added_machno").val(sessionStorage.getItem("added_machno"));
                $("#Added_matetype").val(sessionStorage.getItem("added_matetype"));
                $("#Added_matecode").val(sessionStorage.getItem("added_matecode"));
                $("#Added_matename").val(sessionStorage.getItem("added_matename"));
                $("#Added_intime").val(sessionStorage.getItem("added_intime"));
                $("#Added_indat").val(sessionStorage.getItem("added_indat"));
                $("#Added_usrno").val(sessionStorage.getItem("added_usrno"));
                $("#Added_state").val(sessionStorage.getItem("added_state"));
            });
            sessionStorage.setItem("added_machno", machno);
            sessionStorage.setItem("added_matetype", matetype);
            sessionStorage.setItem("added_matecode", matecode);
            sessionStorage.setItem("added_matename", matename);
            sessionStorage.setItem("added_intime", intime);
            sessionStorage.setItem("added_indat", indat);
            sessionStorage.setItem("added_usrno", usrno);
            sessionStorage.setItem("added_state", state);

        },
        failure: function (response) {
            $.alert(response.d);
            alertify.error('Failed to add user!', 1, function () {
                alert(response.d);
            });
        }
    });
}

function GetAllData() {
    $("#btnOkFind").fadeIn(200);
    $("#findMachno").val('');
    $("#findMateType").val('');
    $("#findMateCode").val('');
    $("#findMateName").val('');
    $("#inserted_tittle").fadeOut(200);
    $("#table_inserted").fadeOut(200);
    $("#table_rubmats").fadeIn(200);
    $("#table_add").fadeOut(200);
    $("#btnOK").fadeOut(200);
    $("#insertTitle").fadeOut(200);
    $("#rows_search").fadeIn(200);
    var ID = $("#id").val();
    var Machno = $("#machno").val();
    var Matetype = $("#matetype").val();
    var Matecode = $("#matecode").val();
    var Matename = $("#matename").val();
    var Intime = $("#intime").val();
    var Indat = $("#indat").val();
    var Usrno = $("#usrno").val();
    var State = $("#state").val();
    $.ajax({
        url: "/getAllData",
        type: "GET",
        data: {
            "id": ID,
            "machno": Machno,
            "matetype": Matetype,
            "matecode": Matecode,
            "matename": Matename,
            "intime": Intime,
            "indat": Indat,
            "usrno": Usrno,
            "state": State
        },
        success: function (response) {
            alertify.notify('Data filtered!.', 'success', 0.5, function () {
                console.log('Success');
            });
            // Xóa toàn bộ dòng trong table và thêm lại dòng mới
            $("#table_rubmats").find("tr:gt(0)").remove();
            $(response).each(function (index, value) {
                var record =
                    "<tr><td>" +
                    (index + 1) +
                    "</td><td>" +
                    value.Machno +
                    "</td><td>" +
                    value.Matetype +
                    "</td><td>" +
                    value.Matecode +
                    "</td><td>" +
                    value.Matename +
                    "</td><td>" +
                    value.Intime +
                    "</td><td>" +
                    value.Indat +
                    "</td><td>" +
                    value.Usrno +
                    "</td><td>" +
                    value.State +
                    "</td><td>" +
                    "</td><tr>";
                $("table").append(record);
            });
        },
        failure: function (response) {
            $.alert(response.d);
        }
    });
}

function DeleteRubmats() {
    $("table").on("click", "#btnDelete", function () {
        var currentRow = $(this).closest("tr");
        var id = currentRow.find("td:eq(1)").text();
        $.ajax({
            url: "/deleteRubmats",
            type: "POST",
            data: {
                'id': id
            },
            success: function (response) {
                $(response).each(function (index, value) {
                    var record =
                        "<tbody><tr><td>" +
                        (index + 1) +
                        "</td><td>" +
                        value.ID +
                        "</td><td>" +
                        value.Machno +
                        "</td><td>" +
                        value.Matetype +
                        "</td><td>" +
                        value.Matecode +
                        "</td><td>" +
                        value.Matename +
                        "</td><td>" +
                        value.Intime +
                        "</td><td>" +
                        value.Indat +
                        "</td><td>" +
                        value.Usrno +
                        "</td><td>" +
                        value.State +
                        "</td><td> " +
                        $("table").append(record);
                });
            },
            failure: function (response) {
                $.alert(response.d);
            }
        });
        $(this).closest('tr').remove();
    });
    alertify.error('Deleted successful!', 'success', 5,
        function () {});
}

function AddRubmats() {
    $('#btnUpdate').fadeOut();
    $("#insertTitle").fadeIn(200);
    $("#table_add").fadeIn(200);
    $('#table_rubmats').fadeOut(200);
    $("#btnOK").fadeIn(200);
    $("#rows_search").fadeIn(200);
    $("#btnOkFind").fadeIn(200);
    $("#machno").val('');
    $("#matetype").val('');
    $("#matename").val('');
    $("#matecode").val('');
    $('#findMachno').val('');
    $('#findMateType').val('');
    $('#findMateCode').val('');
    $('#findMateName').val('');
    $('#table_add').fadeIn(200);
    GetDateTime();
    $('#modifield_Title').fadeOut(200);
    $('#table_modifield').fadeOut(200);
    $('#btnInsert').fadeIn();

}

function Logout() {
    localStorage.clear();
    window.location.assign("/");
}

function CallPrint() {
    $("button").hide();
    $("a").hide();
    $("#col-action").hide();
    var prtContent = document.getElementById("print");
    var WinPrint = window.open('', '',
        'left= 0, top= 0, width= 1, height= 1, toolbar= 0, scrollbars= 0, status= 0');
    WinPrint.document.write(prtContent.innerHTML);
    WinPrint.document.close();
    WinPrint.focus();
    WinPrint.print();
    WinPrint.close();
}

function loadPage(page) {
    $("#divMain").load(page);
}

function SelectButton() {
    $('#insertTitle').fadeOut(200);
    $('#table_add').fadeOut(200);
    $('#btnOkFind').fadeIn(200);
    $("#btnOK").fadeIn(200);
    $('#rows_search').fadeIn(200);
    $('#table_inserted').fadeOut(200);
    $('#inserted_tittle').fadeOut(200);
    $("#table_added").fadeOut(200);
    $("#addedTittle").fadeOut(200);
    $("#table_rubmats").fadeOut(200);
    $("#btnUpdate").fadeOut();

    $('#findMachno').val('');
    $('#findMateType').val('');
    $('#findMateCode').val('');
    $('#findMateName').val('');
    $('#btnInsert').fadeOut();
}


function GetDateTime() {
    var dateTime = new Date($.now());
    var hours = (dateTime.getHours() < 10 ? "0" : "") + dateTime.getHours();
    var minutes = (dateTime.getMinutes() < 10 ? "0" : "") + dateTime.getMinutes();
    var seconds = (dateTime.getSeconds() < 10 ? "0" : "") + dateTime.getSeconds();

    var years = dateTime.getFullYear();
    var months = (dateTime.getMonth() < 10 ? "0" : "") + (parseInt(dateTime.getMonth()) + 1);
    var days = (dateTime.getDate() < 10 ? "0" : "") + dateTime.getDate();

    var indat = years + months + days;
    var intime = hours + ":" + minutes + ":" + seconds;
    $("#intime").val(intime);
    $("#indat").val(indat);
    $("#usrno").val(localStorage.getItem("username"));
    $("#state").val("0");
    $("#modifield_intime").val(intime);
    $("#modifield_indat").val(indat);
    localStorage.setItem("intime", intime);
    localStorage.setItem("indat", indat);
}

function GetModified() {

    //$("#table_modifield").fadeIn(200);
    //$("#modifield_Title").fadeIn(200);
    GetDateTime();
    $("#btnUpdate").fadeIn();
    $("#btnOkFind").fadeOut();
    var ID = "";
    var machno = "";
    var matetype = "";
    var matecode = "";
    var matename = "";
    var intime = "";
    var indat = "";
    var usrno = localStorage.getItem("username");
    var state = "";
    $("table").on("click", "#btnModifield", function () {
        var currentRow = $(this).closest("tr");
        //var ID = $("#mongoID").val();
        ID = currentRow.find("td:eq(1)").text();
        machno = currentRow.find("td:eq(2)").text();
        matetype = currentRow.find("td:eq(3)").text();
        matecode = currentRow.find("td:eq(4)").text();
        matename = currentRow.find("td:eq(5)").text();
        // intime = currentRow.find("td:eq(6)").text();
        // indat = currentRow.find("td:eq(7)").text();
        intime = localStorage.getItem("intime");
        indat = localStorage.getItem("indat");
        //indat = currentRow.find("td:eq(7)").text();
        usrno = localStorage.getItem("username");
        state = currentRow.find("td:eq(9)").text();

        sessionStorage.setItem("mongoID", ID);
        $("#findMachno").val(machno);
        $("#findMateType").val(matetype);
        $("#findMateCode").val(matecode);
        $("#findMateName").val(matename);
        // sessionStorage.setItem("intime", intime);
        // sessionStorage.setItem("indat", indat);
        //sessionStorage.setItem("usrno", usrno);
        //sessionStorage.setItem("state", state);
        //$("#modifield_usrno").val(usrno);
        $("#modifield_state").val(state);
        return true;
    });
}

function UpdateModifield() {

    var updateID = sessionStorage.getItem("mongoID");
    var machno = $("#findMachno").val();
    var matetype = $("#findMateType").val();
    var matecode = $("#findMateCode").val();
    var matename = $("#findMateName").val();
    var intime = localStorage.getItem("intime");
    var indat = localStorage.getItem("indat");
    var usrno = localStorage.getItem("username");
    var state = "0";
    $.ajax({
        url: "/updateRubmat",
        type: "POST",
        data: {
            'updateID': updateID,
            'machno': machno,
            'matetype': matetype,
            'matecode': matecode,
            'matename': matename,
            'intime': intime,
            'indat': indat,
            'usrno': usrno,
            'state': state
        },
        success: function (response) {

            alertify.notify('Update successful!', 'success', 1, function () {
                //$("#updateModal").load();
                $("#findMachno").val('');
                $("#findMateType").val('');
                $("#findMateCode").val('');
                $("#findMateName").val('');
                GetRubmats();
            });
        },
        failure: function (response) {
            //$.alert(response.d);
            var notification = alertify.notify('Update failed', 'danger', 1, function () {
                console.log('dismissed');
            });
        }
    });
}