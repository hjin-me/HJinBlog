/**
 * Created by hjin on 3/1/15.
 */
(function () {
  var detect = document.createElement('detect');
  detect.style.fontSize = '1vw';
  if(detect.style.fontSize !== '1vw') {
    onChangeWidth();
    window.onresize = onChangeWidth;
    window.addEventListener('orientationchange', onChangeWidth);
  }
  function onChangeWidth() {
    document.documentElement.style.fontSize = (window.innerWidth) + 'px';
  }
}());