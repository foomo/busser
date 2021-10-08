/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_busser_service from './vo-service'; // ./client-ts/src/client-service.ts to ./client-ts/src/vo-service.ts
import * as github_com_foomo_busser_table from './vo-table'; // ./client-ts/src/client-service.ts to ./client-ts/src/vo-table.ts
import * as github_com_foomo_busser_table_validation from './vo-validation'; // ./client-ts/src/client-service.ts to ./client-ts/src/vo-validation.ts

export class ServiceClient {
	public static defaultEndpoint = "/services/busser";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async commit(id:github_com_foomo_busser_table.ID, version:github_com_foomo_busser_table.Version):Promise<github_com_foomo_busser_service.ErrorCommit> {
		return (await this.transport<{0:github_com_foomo_busser_service.ErrorCommit}>("Commit", [id, version]))[0]
	}
	async delete(id:github_com_foomo_busser_table.ID, versions:github_com_foomo_busser_table.Version[]):Promise<github_com_foomo_busser_service.ErrorDelete> {
		return (await this.transport<{0:github_com_foomo_busser_service.ErrorDelete}>("Delete", [id, versions]))[0]
	}
	async getCommitted(id:github_com_foomo_busser_table.ID):Promise<{t:github_com_foomo_busser_table.Table; vt:github_com_foomo_busser_table_validation.Table; err:github_com_foomo_busser_service.ErrorGet}> {
		let response = await this.transport<{0:github_com_foomo_busser_table.Table; 1:github_com_foomo_busser_table_validation.Table; 2:github_com_foomo_busser_service.ErrorGet}>("GetCommitted", [id])
		let responseObject = {t : response[0], vt : response[1], err : response[2]};
		return responseObject;
	}
	async getVersion(id:github_com_foomo_busser_table.ID, version:github_com_foomo_busser_table.Version):Promise<{t:github_com_foomo_busser_table.Table; vt:github_com_foomo_busser_table_validation.Table; err:github_com_foomo_busser_service.ErrorGet}> {
		let response = await this.transport<{0:github_com_foomo_busser_table.Table; 1:github_com_foomo_busser_table_validation.Table; 2:github_com_foomo_busser_service.ErrorGet}>("GetVersion", [id, version])
		let responseObject = {t : response[0], vt : response[1], err : response[2]};
		return responseObject;
	}
	async list():Promise<{ret:github_com_foomo_busser_table.Map; ret_1:github_com_foomo_busser_service.ErrorGet}> {
		let response = await this.transport<{0:github_com_foomo_busser_table.Map; 1:github_com_foomo_busser_service.ErrorGet}>("List", [])
		let responseObject = {ret : response[0], ret_1 : response[1]};
		return responseObject;
	}
	async validate(id:github_com_foomo_busser_table.ID):Promise<{t:github_com_foomo_busser_table.Table; vt:github_com_foomo_busser_table_validation.Table; err:github_com_foomo_busser_service.ErrorValidation}> {
		let response = await this.transport<{0:github_com_foomo_busser_table.Table; 1:github_com_foomo_busser_table_validation.Table; 2:github_com_foomo_busser_service.ErrorValidation}>("Validate", [id])
		let responseObject = {t : response[0], vt : response[1], err : response[2]};
		return responseObject;
	}
}