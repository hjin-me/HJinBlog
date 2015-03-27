/**
 * Created by hjin on 2/28/15.
 */
(function (document) {
  var e = document.createElement("detect"), l = !1, t = !1;
  e.style.display = "flex";
  if ("flex" === e.style.display) {
    l = true;
  }
  e.style.display = "-webkit-flex";
  if ("-webkit-flex" === e.style.display) {
    l = true
  }
  e.style.display = "-webkit-box";
  if ("-webkit-box" === e.style.display) {
    t = true;
  }
  if (l) {
    document.documentElement.className += " flex_support";
  } else if (t) {
    document.documentElement.className += " flex_old_support";
  } else {
    document.documentElement.className += " flex_unsupported";
  }

}(document));