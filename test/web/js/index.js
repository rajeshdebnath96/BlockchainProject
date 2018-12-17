$(document).ready(function() {

	// Add a given number of units to the account value
	$("#submit").click(function() {
		var sender = $("#sender").val();
		var receiver = $("#receiver").val();
		var txAmount = $("#txAmount").val();
		var pm = $("#pm").val();
		var txNo = $("#txNo").val();
		var txDate = $('#txDate').val();
		console.log("Sender : " + sender+" "+receiver+" "+txAmount+" "+pm+" "+txNo+" "+txDate);

		$.ajax({
		    beforeSend: function(xhrObj){
		        xhrObj.setRequestHeader("Content-Type","application/json");
		        xhrObj.setRequestHeader("Accept","application/json");
		    },
		    type: "POST",
		    url: 'http://localhost:3001/submit',
			data: JSON.stringify({ "sender": sender,"receiver":receiver,"txAmount":txAmount,"pm":pm,"txNo":txNo,"txDate":txDate }),
		    dataType: "json",
		    success: function(result) {
		       console.log("Status: " + result.status);
		       $("#submitResult").text("Status: Your Transaction id :" + result.status)
		    }
		});
	});
	// Query the current value of the account
	
	$("#viewTx").click(function() {
		var viewTx = $("#viewTxID").val();
		//var viewTx = $("input[name=viewTxID]").val();
		//alert(viewTx1);
		var data = {
			"viewTx" : viewTx
		};
		console.log("viewTxID = " + viewTx);
		$.ajax({
			/*beforeSend: function(xhrObj){
		        xhrObj.setRequestHeader("Content-Type","application/json");
		        xhrObj.setRequestHeader("Accept","application/json");
		    },*/
		    type: "POST",
		    url: 'http://localhost:3001/viewTx',
			// data: JSON.stringify({ "viewTx": viewTx,"viewTx1": viewTx }),
		    data : data,
		    dataType: "json",
			success: function(data, status) {
				console.log(data);
				$("#viewTxResult").text("Status: " + data.status)
				//document.getElementById("accountValue").innerHTML = "Account ($): " + data.value;
			}
		});
	});
	$("#verifyTx").click(function() {
		var verifyTx = $("#verifyTxID").val();
		var verifyTxStatus = $("#verifyTxStatus").val();
		console.log("verifyTx = " + verifyTx+" "+verifyTxStatus);

		$.ajax({
		    beforeSend: function(xhrObj){
		        xhrObj.setRequestHeader("Content-Type","application/json");
		        xhrObj.setRequestHeader("Accept","application/json");
		    },
		    type: "POST",
		    url: 'http://localhost:3001/verifyTx',
			data: JSON.stringify({ "verifyTx": verifyTx,"verifyTxStatus":verifyTxStatus }),
		    dataType: "json",
		    success: function(result) {
		       console.log("Status: " + result.status);
				$("#verifyTxResult").text("Status: " + result.status)
		    }
		});
	});
	$("#viewMyTx").click(function() {
		var viewMyTxID = $("#viewMyTxID").val();
		console.log("viewMyTx = " + viewMyTxID);

		var data = { 
						"viewMyTx": viewMyTxID,
						"viewTx1": "Rajesh" 
					};

		$.ajax({
		    
		    type: "POST",
		    url: 'http://localhost:3001/viewMyTx',
			// data: JSON.stringify(data),
			data: data,
		    dataType: "json",
		    success: function(result) {
		       console.log("Status: " + result.status);
				//$("#viewMyTxResult").text("Status: " + result.status);
				/*var res = result.status.split("&");
				var index = res.indexOf("$$$$$$$");
				var index_valid,valid_status;
				var str = "<table class=\"table table-hover\">";
				str += "<thead><tr>"
				str += "<th>Transaction ID</th><th>Sender</th><th>Receiver</th><th>Transaction Amount</th>";
				str += "<th>Payment Media</th><th>Transaction No</th><th>Verification Status</th><tr></thead>";
				for(var i=0;(i+1)*6<index;i++){
					str += "<tr ";
					index_valid = res.indexOf(res[i*6],index);
					if (res[index_valid+1] == 0 ) {str+="class=\"info\">" 
						valid_status ="Not Verified";
					}
					else if (res[index_valid+1] == 1) {str+="class=\"success\">"
						valid_status="Correct (verified)";
					}
					else if (res[index_valid+1] == 2) {str+="class=\"danger\">"
						valid_status="Incorrect (verified)";
					}
					for(var j=0;j<6;j++ ){
						str+="<td>";
						str += res[i*6+j];
						str += "</td>"
					}
					str += "<td>";
					
					console.log(res,index_valid,index);
					str+=valid_status
					//str += res[index_valid+1];
					str += "</td></tr>"
				}
				str += "</table>"; */
				var res = result.status.split("&");
				var valid_str1, valid_str2;
				var index = res.indexOf("$$$$$$$")
				var index_valid; 

				var c = 7;
				var str = "<table class=\"table table-hover\">";
				str += "<thead><tr>"
				str += "<th>Transaction ID</th><th>Sender</th><th>Receiver</th><th>Transaction Amount</th>";
				str += "<th>Payment Media</th><th>Transaction No</th><th>Transaction Date</th><th>Verification Status</th><tr></thead>";
				for(var i=0;(i+1)*c<index;i++){
					str += "<tr class =";
					index_valid = res.indexOf(res[i*c],index);
					console.log(res,index_valid,index);
					if (res[index_valid+1] == 0 ) {valid_str1 ="active"; valid_str2="Not Verified";}
					else if (res[index_valid+1] == 1) {valid_str1 ="success"; valid_str2="Correct (verified)";}
					else if (res[index_valid+1] == 2) {valid_str1 ="danger"; valid_str2="Incorrect (verified)";} 
					str += valid_str1 + ">"
					for(var j=0;j<c;j++ ){
						str+="<td>";
						str += res[i*c+j];
						str += "</td>"
					}
					str += "<td>";
					/*index_valid = res.indexOf(res[i*c],index);
					console.log(res,index_valid,index);
					if (res[index_valid+1] == 0 ) valid_str1 ="active" str+="Not Verified";
					else if (res[index_valid+1] == 1) valid_str1 ="success" str+="Correct (verified)";
					else if (res[index_valid+1] == 2) valid_str1 ="danger" str+="Incorrect (verified)"; */
					str += valid_str2;
					//str += res[index_valid+1];
					str += "</td></tr>"
				}
				str += "</table>";
				$("#viewMyTxResult").html(str);
		    }
		});
	});
	$("#viewAllTx").click(function() {
		//var amount = $("input").val();
		console.log("All Transactions = ");

		$.ajax({
		    beforeSend: function(xhrObj){
		        xhrObj.setRequestHeader("Content-Type","application/json");
		        xhrObj.setRequestHeader("Accept","application/json");
		    },
		    type: "GET",
		    url: 'http://localhost:3001/viewAllTx',
			//data: JSON.stringify({ "amount": amount }),
		    dataType: "json",
		    success: function(result) {
		       console.log("Status: " + result.status);
				//$("#viewAllTxResult").text("Status: " + result.status);
				var res = result.status.split("&");
				var valid_str1, valid_str2;
				var index = res.indexOf("$$$$$$$")
				var index_valid; 

				var c = 7;
				var str = "<table class=\"table table-hover\">";
				str += "<thead><tr>"
				str += "<th>Transaction ID</th><th>Sender</th><th>Receiver</th><th>Transaction Amount</th>";
				str += "<th>Payment Media</th><th>Transaction No</th><th>Transaction Date</th><th>Verification Status</th><tr></thead>";
				for(var i=0;(i+1)*c<index;i++){
					str += "<tr class =";
					index_valid = res.indexOf(res[i*c],index);
					console.log(res,index_valid,index);
					if (res[index_valid+1] == 0 ) {valid_str1 ="active"; valid_str2="Not Verified";}
					else if (res[index_valid+1] == 1) {valid_str1 ="success"; valid_str2="Correct (verified)";}
					else if (res[index_valid+1] == 2) {valid_str1 ="danger"; valid_str2="Incorrect (verified)";} 
					str += valid_str1 + ">"
					for(var j=0;j<c;j++ ){
						str+="<td>";
						str += res[i*c+j];
						str += "</td>"
					}
					str += "<td>";
					/*index_valid = res.indexOf(res[i*c],index);
					console.log(res,index_valid,index);
					if (res[index_valid+1] == 0 ) valid_str1 ="active" str+="Not Verified";
					else if (res[index_valid+1] == 1) valid_str1 ="success" str+="Correct (verified)";
					else if (res[index_valid+1] == 2) valid_str1 ="danger" str+="Incorrect (verified)"; */
					str += valid_str2;
					//str += res[index_valid+1];
					str += "</td></tr>"
				}
				str += "</table>";
				$("#viewAllTxResult").html(str);
		    }
		});
	});
	$("#submit_link").click(function(){
		//alert("gfhgfhhgf");
		$("#submit").load("html/submit.html");
	});
	$("#view_link").click(function(){
		//alert("gfhgfhhgf");
		$("#submit").load("html/view.html");
	});
	$("#verify_link").click(function(){
		//alert("gfhgfhhgf");
		$("#submit").load("html/verify.html");
	});
	$("#mytx_link").click(function(){
		//alert("gfhgfhhgf");
		$("#submit").load("html/my_tx.html");
	});
	$("#alltx_link").click(function(){
		//alert("gfhgfhhgf");
		$("#submit").load("html/all_tx.html");
	});
});