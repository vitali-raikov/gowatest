require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
let TinyDatePicker = require("tiny-date-picker")

$(() => {
    TinyDatePicker('.datepicker');
});
