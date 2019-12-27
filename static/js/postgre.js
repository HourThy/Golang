function GetPostGre() {
    // var ID = "";
    // var Username = "";
    // var Password = "";
    $.ajax({
        url: "/connectToPostgreSQL",
        type: "GET",
        // data: {
        //     "id": ID,
        //     "username": Username,
        //     "password": Password
        // },
        contentType: "application/json",
        success: function (response) {
            $("#table_postgre").find("tr:gt(0)").remove();

            alertify.notify('Data filtered!.', 'success', 0.5, function () {
                $(response).each(function (index, value) {
                    var record =
                        "<tr><td>" +
                        (index + 1) +
                        "</td><td>" +
                        value.username +
                        "</td><td>" +
                        value.password +
                        "</td><tr>";
                    record = record.substring(0, record.length - 9);
                    $("table").append(record);
                });
            });

        },
        failure: function (response) {
            alert("Get data error!");
            $("#table_postgre").hide();
        }
    });
}