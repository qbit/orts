function totalCounts(a) {
    var counts = {}, i, l, ret = [];
    for (i = 0, l = a.length; i < l; i++) {
	var name = a[i].innerText;
	if (! counts[name]) {
	    counts[name] = 1;
	} else {
	    counts[name] += 1;
	}
    }				    

    for (var c in counts) {
	ret.push([c, counts[c]]);
    }
    return ret;
}
function countsByArch() {
    var a = document.getElementsByClassName("arch");
    return totalCounts(a);
}

function countsByMedia() {
    var a = document.getElementsByClassName("boot-media");
    return totalCounts(a);
}

function countsByInstType() {
    var a = document.getElementsByClassName("install-method");
    return totalCounts(a);
}

function countsByStatus() {
    var a = document.getElementsByClassName("success-fail");
    return totalCounts(a);
}

function chartDefault(name, agg) {

    var defaultChart = {
	chart: {
	    plotBackgroundColor: null,
	    plotBorderWidth: 0,
	    plotShadow: false
	},
	title: {
	    text: '',
	    align: 'center',
	    verticalAlign: 'middle',
	    y: 40
	},
	tooltip: {
	    pointFormat: '<b>{point.percentage:.1f}%</b>'
	},
	plotOptions: {
	    pie: {
		dataLabels: {
		    enabled: true,
		    distance: 20,
		    style: {
			fontWeight: 'bold',
			color: 'black'
		    }
		},
		startAngle: -90,
		endAngle: 90,
		center: ['50%', '75%']
	    }
	},
	series: [{
	    type: 'pie',
	    name: name,
	    innerSize: '50%',
	    data: agg()
	}]
    };

    return defaultChart;
}

$(function () {
    $('#tests-by-arch').highcharts(chartDefault('Tests by arch', countsByArch));
    $('#tests-by-media').highcharts(chartDefault('Tests by media', countsByMedia));
    $('#tests-ins-vs-up').highcharts(chartDefault('Install vs Upgrade', countsByInstType));
    $('#tests-suc-vs-fail').highcharts(chartDefault('Success vs Fail', countsByStatus));
});