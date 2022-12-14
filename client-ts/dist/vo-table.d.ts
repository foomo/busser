import * as github_com_foomo_busser_table from './vo-table';
export type ColumnName = string;
export type ID = string;
export type List = Array<github_com_foomo_busser_table.TableSummary>;
export type Map = Record<github_com_foomo_busser_table.ID, github_com_foomo_busser_table.List | null>;
export type Row = Record<github_com_foomo_busser_table.ColumnName, string>;
export type Rows = Array<github_com_foomo_busser_table.Row | null>;
export interface Table {
    id: github_com_foomo_busser_table.ID;
    version: github_com_foomo_busser_table.Version;
    timestamp: number;
    rows: github_com_foomo_busser_table.Rows | null;
    readErrors: Array<string> | null;
}
export interface TableSummary {
    id: github_com_foomo_busser_table.ID;
    timestamp: number;
    version: github_com_foomo_busser_table.Version;
    valid: boolean;
    committed: boolean;
}
export type Version = string;
