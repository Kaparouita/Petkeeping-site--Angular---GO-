export class LoginResp{
    username: string;
    password: string;
    user_type: string;
    user_id: number;

    constructor(
        username: string,
        password: string,
        user_type: string,
        user_id: number,
    ) {
        this.username = username;
        this.password = password;
        this.user_type = user_type;
        this.user_id = user_id;
    }
}