const { readFileSync, writeFileSync } = require('fs');
module.exports = function () {
    var processed = readFileSync('../data/userData.json', 'utf8');
    processed = processed.substring(0, processed.length - 1);
    processed += '}';
    return processed;
}