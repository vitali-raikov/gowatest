require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

let TinyDatePicker = require("tiny-date-picker");

let displayFormat = 'MM/DD/YYYY';
import moment from 'moment/src/moment'


$(() => {
    TinyDatePicker('.datepicker', {
        mode: 'dp-below',
        format(date) {
            return moment(date).format(displayFormat);
        },
        parse(str) {
            return moment(str, displayFormat).toDate();
        },
        max: moment().subtract(18, 'years').format(displayFormat),
        min: moment().subtract(60, 'years').format(displayFormat)
    });
});
