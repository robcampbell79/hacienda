create table if not exists hacienda_projects (
	project_id int auto_increment primary key,
    project_name varchar(255) not null,
    num_of_task int null,
    num_of_members int null,
    created timestamp default current_timestamp);