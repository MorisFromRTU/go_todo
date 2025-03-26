export interface AuthForm {
    username: string;
    password: string;
}

export interface Task {
    id: number;
    title: string;
    done: boolean; 
    userId: number;
}

export interface NewTask {
    title: string;
    done: boolean; 
    userId: number;
}

export interface AuthFormProps {
    title: string;
    buttonText: string;
    onSubmit: (username: string, password: string) => void;
  }