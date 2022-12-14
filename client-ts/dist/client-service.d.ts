import * as github_com_foomo_busser_service from './vo-service';
import * as github_com_foomo_busser_table from './vo-table';
import * as github_com_foomo_busser_table_validation from './vo-validation';
export declare class ServiceClient {
    transport: <T>(method: string, data?: any[]) => Promise<T>;
    static defaultEndpoint: string;
    constructor(transport: <T>(method: string, data?: any[]) => Promise<T>);
    commit(id: github_com_foomo_busser_table.ID, version: github_com_foomo_busser_table.Version): Promise<github_com_foomo_busser_service.ErrorCommit | null>;
    delete(id: github_com_foomo_busser_table.ID, versions: Array<github_com_foomo_busser_table.Version> | null): Promise<github_com_foomo_busser_service.ErrorDelete | null>;
    getCommitted(id: github_com_foomo_busser_table.ID): Promise<{
        t: github_com_foomo_busser_table.Table | null;
        vt: github_com_foomo_busser_table_validation.Table | null;
        err: github_com_foomo_busser_service.ErrorGet | null;
    }>;
    getVersion(id: github_com_foomo_busser_table.ID, version: github_com_foomo_busser_table.Version): Promise<{
        t: github_com_foomo_busser_table.Table | null;
        vt: github_com_foomo_busser_table_validation.Table | null;
        err: github_com_foomo_busser_service.ErrorGet | null;
    }>;
    list(): Promise<{
        ret: github_com_foomo_busser_table.Map;
        ret_1: github_com_foomo_busser_service.ErrorGet | null;
    }>;
    validate(id: github_com_foomo_busser_table.ID): Promise<{
        t: github_com_foomo_busser_table.Table | null;
        vt: github_com_foomo_busser_table_validation.Table | null;
        err: github_com_foomo_busser_service.ErrorValidation | null;
    }>;
}
