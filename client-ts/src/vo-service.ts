/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_busser_service from './vo-service'; // ./client-ts/src/vo-service.ts to ./client-ts/src/vo-service.ts
import * as github_com_foomo_busser_table from './vo-table'; // ./client-ts/src/vo-service.ts to ./client-ts/src/vo-table.ts
import * as github_com_foomo_busser_table_validation from './vo-validation'; // ./client-ts/src/vo-service.ts to ./client-ts/src/vo-validation.ts
// github.com/foomo/busser/service.ErrorCommit
export enum ErrorCommit {
	CouldNotCommit = "COULD_NOT_COMMIT",
}
// github.com/foomo/busser/service.ErrorDelete
export enum ErrorDelete {
	CouldNotDeleteVersion = "COULD_NOT_DELETE_VERSION",
}
// github.com/foomo/busser/service.ErrorGet
export enum ErrorGet {
	CouldNotLoadTableFromStore = "COULD_NOT_LOAD_TABLE_FROM_STORE",
}
// github.com/foomo/busser/service.ErrorValidation
export enum ErrorValidation {
	CouldNotValidate = "COULD_NOT_LOAD_TABLE_FROM_STORE",
}
// end of common js