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