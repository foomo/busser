import * as github_com_foomo_busser_table from './vo-table';
import * as github_com_foomo_busser_table_validation from './vo-validation';
export interface Cell {
    valid: boolean;
    feedback: github_com_foomo_busser_table_validation.Feedback | null;
}
export type Feedback = Array<github_com_foomo_busser_table_validation.FeedbackEntry>;
export interface FeedbackEntry {
    level: github_com_foomo_busser_table_validation.FeedbackLevel;
    msg: string;
}
export declare enum FeedbackLevel {
    Error = "error",
    Valid = "valid",
    Warning = "warning"
}
export interface Row {
    valid: boolean;
    cells: Record<github_com_foomo_busser_table.ColumnName, github_com_foomo_busser_table_validation.Cell> | null;
    feedback: github_com_foomo_busser_table_validation.Feedback | null;
}
export type Rows = Array<github_com_foomo_busser_table_validation.Row | null>;
export interface Table {
    valid: boolean;
    rows: github_com_foomo_busser_table_validation.Rows | null;
    feedback: github_com_foomo_busser_table_validation.Feedback | null;
}
