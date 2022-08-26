## Api simples com PostGres

<hr>
<h3>Para criar o banco use os comandos</h3>
<br>
<strong>docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5</strong>

<br>
<strong>docker exec -it api-todo psql -U postgres</strong>

<br>
<strong>create database api_todo;</strong>

<br>
<strong>create user user_todo;</strong>

<br>
<strong>alter user user_todo with encrypted password '1122';</strong>

<br>
<strong>grant all privileges on database api_todo to user_todo;</strong>

<br>
<h4>Crie a tabela<h4>


<strong>\c api_todo;</strong>

<strong>create table todos (id serial primary key, title varchar, description text, done bool);</strong>

verifique a tabela com o comando <strong>\dt</strong>

Forneça as permissões com o comando 
<br>
<strong>grant all privileges on all tables in schema public to user_todo;</strong><br>
<strong>grant all privileges on all sequences in schema public to user_todo;</strong><br>
<br>

<hr>
<h3>Para executar a Api use o comando</h3>
<br>
<strong>go run main.go</strong>
