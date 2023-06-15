import { reactive } from 'vue';

type userKeys = 'id' | 'firstname' | 'lastname' | 'phone' | 'email' | 'role' | 'username';
type userData = { [key: string]: string | number; };

export const user = reactive({
  id: '',
  firstname: '',
  lastname: '',
  phone: '',
  email: '',
  role: '0',
  username: '',
  setAll (data: userData) {
    (Object.keys(data) as userKeys[]).forEach((key) => {
      this.set(key, data[key]);
    });
  },
  set (key: userKeys, value: any) {
    this[key] = value; sessionStorage.setItem(`user_${key}`, value);
  },
  get (key: userKeys) {
    if (!this[key] && sessionStorage.getItem(`user_${key}`)) {
      this[key] = sessionStorage.getItem(`user_${key}`) as any;
    }
    return this[key];
  },
});