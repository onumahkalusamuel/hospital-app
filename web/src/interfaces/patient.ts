export interface Patient {
  id: string;
  create_at: Date,
  card_no: string | number;
  title: string;
  firstname: string;
  middlename: string;
  lastname: string;
  phone: string;
  current_appointment: string;
  sex: string;
}