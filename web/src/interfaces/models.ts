import { PatientHistoryTypes, RoleTypes } from '.';

export interface Patient {
  id: string;
  created_at: Date,
  card_no: string | number;
  title: string;
  firstname: string;
  middlename: string;
  lastname: string;
  phone: string;
  email: string;
  dob: string;
  current_appointment: string;
  current_status: string;
  sex: string;
  address: string;
  occupation: string;
  marital_status: string;
  next_of_kin?: NextOfKin;
  patient_history?: PatientHistory[];
  invoices?: Invoice[];
}

export interface NextOfKin {
  id: string;
  created_at: Date,
  title: string;
  firstname: string;
  middlename: string;
  lastname: string;
  phone: string;
  email: string;
  sex: string;
  address: string;
  occupation: string;
  relationship: string;
}

export interface Delivery {
  id: string;
  patient_id: string,
  delivery_date_time: string;
  delivery_mode: string;
  baby_sex: string;
  baby_weight: number;
  condition: string;
  note: string;
  patient?: Patient;
}

export interface Appointment {
  id: string;
  patient_id: string;
  date_time: string;
  type: string;
  note: string;
}

export interface Invoice {
  id: string;
  name: string;
  billing_address: string;
  invoice_number: string;
  invoice_date: string;
  due_date: string;
  details: InvoiceDetails[];
  patient_id: string;
  amount: number;
  balance: number;
  completed: number;
  discount: number;
  payments?: Payment[];
  patient?: Patient;
}

export interface InvoiceDetails {
  description: string;
  qty: number;
  unit: string;
  price: number;
  amount: number;
}

export interface Payment {
  id: string;
  patient_id: string;
  invoice_id: string;
  note: string;
  amount_paid: number;
  balance: number;
  payment_reference: string;
  status: string;
  payment_proof: string;
}

export interface PatientHistory {
  id: string;
  patient_id: string;
  date_time: string;
  type: PatientHistoryTypes;
  details: { [key: string]: any; };
  patient?: Patient;
}

export interface Settings {
  setting: string;
  value: string;
}

export interface Staff {
  id: string;
  created_at?: Date,
  firstname: string;
  middlename: string;
  lastname: string;
  phone: string;
  email: string;
  sex: string;
  role: RoleTypes;
  username: string;
  password?: string;
}


export interface Pagination {
  limit: number;
  page: number;
  rows: any[];
  sort_key: string;
  sort_order: "asc" | "desc";
  total_pages: number;
  total_rows: number;
  query: '';
}