(window.webpackJsonp=window.webpackJsonp||[]).push([[0],[,function(e,n,t){"use strict";
/*
object-assign
(c) Sindre Sorhus
@license MIT
*/var r=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,o=Object.prototype.propertyIsEnumerable;function l(e){if(null==e)throw new TypeError("Object.assign cannot be called with null or undefined");return Object(e)}e.exports=function(){try{if(!Object.assign)return!1;var e=new String("abc");if(e[5]="de","5"===Object.getOwnPropertyNames(e)[0])return!1;for(var n={},t=0;t<10;t++)n["_"+String.fromCharCode(t)]=t;if("0123456789"!==Object.getOwnPropertyNames(n).map((function(e){return n[e]})).join(""))return!1;var r={};return"abcdefghijklmnopqrst".split("").forEach((function(e){r[e]=e})),"abcdefghijklmnopqrst"===Object.keys(Object.assign({},r)).join("")}catch(e){return!1}}()?Object.assign:function(e,n){for(var t,u,i=l(e),c=1;c<arguments.length;c++){for(var s in t=Object(arguments[c]))a.call(t,s)&&(i[s]=t[s]);if(r){u=r(t);for(var f=0;f<u.length;f++)o.call(t,u[f])&&(i[u[f]]=t[u[f]])}}return i}},,,,function(e,n,t){"use strict";e.exports=t(6)},function(e,n,t){"use strict";
/** @license React v0.18.0
 * scheduler.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var r,a,o,l,u;if(Object.defineProperty(n,"__esModule",{value:!0}),"undefined"==typeof window||"function"!=typeof MessageChannel){var i=null,c=null,s=function(){if(null!==i)try{var e=n.unstable_now();i(!0,e),i=null}catch(e){throw setTimeout(s,0),e}},f=Date.now();n.unstable_now=function(){return Date.now()-f},r=function(e){null!==i?setTimeout(r,0,e):(i=e,setTimeout(s,0))},a=function(e,n){c=setTimeout(e,n)},o=function(){clearTimeout(c)},l=function(){return!1},u=n.unstable_forceFrameRate=function(){}}else{var b=window.performance,m=window.Date,p=window.setTimeout,d=window.clearTimeout;if("undefined"!=typeof console){var y=window.cancelAnimationFrame;"function"!=typeof window.requestAnimationFrame&&console.error("This browser doesn't support requestAnimationFrame. Make sure that you load a polyfill in older browsers. https://fb.me/react-polyfills"),"function"!=typeof y&&console.error("This browser doesn't support cancelAnimationFrame. Make sure that you load a polyfill in older browsers. https://fb.me/react-polyfills")}if("object"==typeof b&&"function"==typeof b.now)n.unstable_now=function(){return b.now()};else{var w=m.now();n.unstable_now=function(){return m.now()-w}}var v=!1,h=null,g=-1,_=5,E=0;l=function(){return n.unstable_now()>=E},u=function(){},n.unstable_forceFrameRate=function(e){0>e||125<e?console.error("forceFrameRate takes a positive int between 0 and 125, forcing framerates higher than 125 fps is not unsupported"):_=0<e?Math.floor(1e3/e):5};var j=new MessageChannel,k=j.port2;j.port1.onmessage=function(){if(null!==h){var e=n.unstable_now();E=e+_;try{h(!0,e)?k.postMessage(null):(v=!1,h=null)}catch(e){throw k.postMessage(null),e}}else v=!1},r=function(e){h=e,v||(v=!0,k.postMessage(null))},a=function(e,t){g=p((function(){e(n.unstable_now())}),t)},o=function(){d(g),g=-1}}function O(e,n){var t=e.length;e.push(n);e:for(;;){var r=Math.floor((t-1)/2),a=e[r];if(!(void 0!==a&&0<S(a,n)))break e;e[r]=n,e[t]=a,t=r}}function T(e){return void 0===(e=e[0])?null:e}function x(e){var n=e[0];if(void 0!==n){var t=e.pop();if(t!==n){e[0]=t;e:for(var r=0,a=e.length;r<a;){var o=2*(r+1)-1,l=e[o],u=o+1,i=e[u];if(void 0!==l&&0>S(l,t))void 0!==i&&0>S(i,l)?(e[r]=i,e[u]=t,r=u):(e[r]=l,e[o]=t,r=o);else{if(!(void 0!==i&&0>S(i,t)))break e;e[r]=i,e[u]=t,r=u}}}return n}return null}function S(e,n){var t=e.sortIndex-n.sortIndex;return 0!==t?t:e.id-n.id}var M=[],P=[],I=1,A=null,C=3,F=!1,N=!1,q=!1;function L(e){for(var n=T(P);null!==n;){if(null===n.callback)x(P);else{if(!(n.startTime<=e))break;x(P),n.sortIndex=n.expirationTime,O(M,n)}n=T(P)}}function U(e){if(q=!1,L(e),!N)if(null!==T(M))N=!0,r(B);else{var n=T(P);null!==n&&a(U,n.startTime-e)}}function B(e,t){N=!1,q&&(q=!1,o()),F=!0;var r=C;try{for(L(t),A=T(M);null!==A&&(!(A.expirationTime>t)||e&&!l());){var u=A.callback;if(null!==u){A.callback=null,C=A.priorityLevel;var i=u(A.expirationTime<=t);t=n.unstable_now(),"function"==typeof i?A.callback=i:A===T(M)&&x(M),L(t)}else x(M);A=T(M)}if(null!==A)var c=!0;else{var s=T(P);null!==s&&a(U,s.startTime-t),c=!1}return c}finally{A=null,C=r,F=!1}}function D(e){switch(e){case 1:return-1;case 2:return 250;case 5:return 1073741823;case 4:return 1e4;default:return 5e3}}var J=u;n.unstable_ImmediatePriority=1,n.unstable_UserBlockingPriority=2,n.unstable_NormalPriority=3,n.unstable_IdlePriority=5,n.unstable_LowPriority=4,n.unstable_runWithPriority=function(e,n){switch(e){case 1:case 2:case 3:case 4:case 5:break;default:e=3}var t=C;C=e;try{return n()}finally{C=t}},n.unstable_next=function(e){switch(C){case 1:case 2:case 3:var n=3;break;default:n=C}var t=C;C=n;try{return e()}finally{C=t}},n.unstable_scheduleCallback=function(e,t,l){var u=n.unstable_now();if("object"==typeof l&&null!==l){var i=l.delay;i="number"==typeof i&&0<i?u+i:u,l="number"==typeof l.timeout?l.timeout:D(e)}else l=D(e),i=u;return e={id:I++,callback:t,priorityLevel:e,startTime:i,expirationTime:l=i+l,sortIndex:-1},i>u?(e.sortIndex=i,O(P,e),null===T(M)&&e===T(P)&&(q?o():q=!0,a(U,i-u))):(e.sortIndex=l,O(M,e),N||F||(N=!0,r(B))),e},n.unstable_cancelCallback=function(e){e.callback=null},n.unstable_wrapCallback=function(e){var n=C;return function(){var t=C;C=n;try{return e.apply(this,arguments)}finally{C=t}}},n.unstable_getCurrentPriorityLevel=function(){return C},n.unstable_shouldYield=function(){var e=n.unstable_now();L(e);var t=T(M);return t!==A&&null!==A&&null!==t&&null!==t.callback&&t.startTime<=e&&t.expirationTime<A.expirationTime||l()},n.unstable_requestPaint=J,n.unstable_continueExecution=function(){N||F||(N=!0,r(B))},n.unstable_pauseExecution=function(){},n.unstable_getFirstCallbackNode=function(){return T(M)},n.unstable_Profiling=null},function(e,n,t){},function(e,n,t){},function(e,n,t){},function(e,n,t){"use strict";t.r(n);var r=t(0),a=t.n(r),o=t(2),l=(t(7),function(){return a.a.createElement("header",{className:"head"},a.a.createElement("img",{src:"https://bitcoin.org/img/icons/logotop.svg?1581019563",alt:"bitcoin"}))});t(8),t(9);function u(e){return function(e){if(Array.isArray(e)){for(var n=0,t=new Array(e.length);n<e.length;n++)t[n]=e[n];return t}}(e)||function(e){if(Symbol.iterator in Object(e)||"[object Arguments]"===Object.prototype.toString.call(e))return Array.from(e)}(e)||function(){throw new TypeError("Invalid attempt to spread non-iterable instance")}()}function i(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){if(!(Symbol.iterator in Object(e)||"[object Arguments]"===Object.prototype.toString.call(e)))return;var t=[],r=!0,a=!1,o=void 0;try{for(var l,u=e[Symbol.iterator]();!(r=(l=u.next()).done)&&(t.push(l.value),!n||t.length!==n);r=!0);}catch(e){a=!0,o=e}finally{try{r||null==u.return||u.return()}finally{if(a)throw o}}return t}(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance")}()}function c(e){return e.toString().replace(/\B(?=(\d{3})+(?!\d))/g,",")}var s=function(){var e=i(Object(r.useState)([]),2),n=e[0],t=e[1],o=new WebSocket("ws://".concat(window.location.host,"/ws"));return Object(r.useEffect)((function(){o.onopen=function(){return console.log("connected")},o.onmessage=function(e){var n=JSON.parse(e.data);t(u(n))},o.onclose=function(){return console.log("disconnected")}}),[o.onclose,o.onmessage,o.onopen]),a.a.createElement("main",null,a.a.createElement("table",null,a.a.createElement("thead",null,a.a.createElement("tr",null,a.a.createElement("th",null,"來源"),a.a.createElement("th",null,"價格"),a.a.createElement("th",null,"24小時交易量"),a.a.createElement("th",null,"24小時漲跌幅度"),a.a.createElement("th",null,"總市值"))),a.a.createElement("tbody",null,n.map((function(e){return a.a.createElement("tr",{key:e.source_name,className:"fadeIn"},a.a.createElement("td",null,e.source_name),a.a.createElement("td",null,"US$".concat(c(Math.round(e.price)))),a.a.createElement("td",null,"US$".concat(c(Math.round(e.volume_24h)))),a.a.createElement("td",null,"".concat(Math.round(100*e.percent_change_24h)/100,"%")),a.a.createElement("td",null,"US$".concat(c(Math.round(e.market_cap)))))})))))},f=function(){return a.a.createElement(a.a.Fragment,null,a.a.createElement(l,null),a.a.createElement(s,null))};Object(o.render)(a.a.createElement(f,null),document.getElementById("app"))}],[[10,1,2]]]);
//# sourceMappingURL=client.695b4bad.chunk.js.map