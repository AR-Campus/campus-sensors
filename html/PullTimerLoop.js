//Timer Loop to pull data every
setInterval(timerLoop, 3000);

function timerLoop()
{
  // console.log("PullTimerLoop-File exec");
  updateWindowsVar = JSON.parse(httpGet("updatewindows"));
  // console.log(updateWindowsVar);
  updateWindows(updateWindowsVar);
  updateTopChartVar = JSON.parse(httpGet("updatetopchart"));
  // updateTopChart(updateTopChartVar)
  updateBottomChartVar = JSON.parse(httpGet("updatebottomchart"));
  // updateBottomChart(updateBottomChartVar)
}


function httpGet(theUrl)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false ); // false for synchronous request
    xmlHttp.send( null );
    return xmlHttp.responseText;
}

function updateWindows(WindowStatus)
{
    statBLi = WindowStatus.BakerStrFensterLi;
    console.log(statBLi);
    statBRe = WindowStatus.BakerStrFensterRe;
    console.log(statBRe);
    statKLi = WindowStatus.KuecheFensterLi;
    console.log(statKLi);
    statKRe = WindowStatus.KuecheFensterRe;
    console.log(statKRe);
    if (statKLi) {
      document.getElementById("KF-Li-Offen").style.opacity = 1;
    } else {
      document.getElementById("KF-Li-Offen").style.opacity = 0;
    }
    if (statKRe) {
      document.getElementById("KF-Re-Offen").style.opacity = 1;
    } else {
      document.getElementById("KF-Re-Offen").style.opacity = 0;
    }
    if (statBLi) {
      document.getElementById("BSF-Li-Offen").style.opacity = 1;
    } else {
      document.getElementById("BSF-Li-Offen").style.opacity = 0;
    }
    if (statBRe) {
      document.getElementById("BSF-Re-Offen").style.opacity = 1;
    } else {
      document.getElementById("BSF-Re-Offen").style.opacity = 0;
    }
    return;
}

function updateTopChart(WindowStatus)
{

}

function updateBottomChart(WindowStatus)
{

}
