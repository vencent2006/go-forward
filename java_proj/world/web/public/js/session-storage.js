SessionStorage = {
    get: function (key) {
        let value = sessionStorage.getItem(key);
        if (value && value.charAt(0) === '{' && value.charAt(value.length - 1) === '}') {
            value = JSON.parse(value);
        }
        return value;
    },
    set: function (key, value) {
        if (typeof value === 'object') {
            value = JSON.stringify(value);
        }
        sessionStorage.setItem(key, value);
    },

    remove: function (key) {
        sessionStorage.removeItem(key);
    },
    clearAll: function () {
        sessionStorage.clear();
    }
};