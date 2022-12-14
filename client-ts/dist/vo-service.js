"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ErrorValidation = exports.ErrorGet = exports.ErrorDelete = exports.ErrorCommit = void 0;
var ErrorCommit;
(function (ErrorCommit) {
    ErrorCommit["CouldNotCommit"] = "COULD_NOT_COMMIT";
})(ErrorCommit = exports.ErrorCommit || (exports.ErrorCommit = {}));
var ErrorDelete;
(function (ErrorDelete) {
    ErrorDelete["CouldNotDeleteVersion"] = "COULD_NOT_DELETE_VERSION";
})(ErrorDelete = exports.ErrorDelete || (exports.ErrorDelete = {}));
var ErrorGet;
(function (ErrorGet) {
    ErrorGet["CouldNotLoadTableFromStore"] = "COULD_NOT_LOAD_TABLE_FROM_STORE";
})(ErrorGet = exports.ErrorGet || (exports.ErrorGet = {}));
var ErrorValidation;
(function (ErrorValidation) {
    ErrorValidation["CouldNotValidate"] = "COULD_NOT_LOAD_TABLE_FROM_STORE";
})(ErrorValidation = exports.ErrorValidation || (exports.ErrorValidation = {}));
