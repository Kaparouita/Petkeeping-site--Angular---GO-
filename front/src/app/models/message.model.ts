export class Message{
    id: number = 0;
    from_name: string = '';
    from_id: number = 0;
    to_name: string = '';
    to_id: number = 0;
    content: string = '';
    date: string = '';

    constructor(
        from_name: string = '',
        to_name: string = '',
        content: string = '',
        date: string = ''
    ) {
        this.from_name = from_name;
        this.to_name = to_name;
        this.content = content;
        this.date = date;
    }
}