// import styles
import '@/assets/sass/main.scss';
//import global libs
import dl from 'lib/js/devLogger';
import jQuery from "jquery";
// set libs to global scope
window.dl = dl;
window.$ = jQuery;
// import global constants
import {isDev} from "lib/js/devLogger";
// set global constants
Object.defineProperty(window, 'isDev', { value: isDev });
// require main js file
require("js/index");
require("js/input-ip-host");