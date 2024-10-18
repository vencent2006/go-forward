package org.example;

import org.apache.calcite.adapter.jdbc.JdbcSchema;
import org.apache.calcite.jdbc.CalciteConnection;
import org.apache.calcite.schema.Schema;
import org.apache.calcite.schema.SchemaPlus;
import org.apache.commons.dbcp2.BasicDataSource;


import java.sql.*;
import java.util.Properties;

public class Main {
    public static void main(String[] args) throws ClassNotFoundException, SQLException {
        System.out.println("Hello world!");
        // check driver exist
        Class.forName("org.apache.calcite.jdbc.Driver");
        Class.forName("com.mysql.cj.jdbc.Driver");

        // the properties for calcite connection
        Properties info = new Properties();
        info.setProperty("lex", "JAVA");
        info.setProperty("remarks","true");
        // SqlParserImpl can analysis sql dialect for sql parse
        info.setProperty("parserFactory","org.apache.calcite.sql.parser.impl.SqlParserImpl#FACTORY");

        // create calcite connection and schema
        Connection connection = DriverManager.getConnection("jdbc:calcite:", info);
        CalciteConnection calciteConnection = connection.unwrap(CalciteConnection.class);
        System.out.println(calciteConnection.getProperties());
        SchemaPlus rootSchema = calciteConnection.getRootSchema();


        // code for mysql datasource
        // MysqlDataSource dataSource = new MysqlDataSource();
        BasicDataSource dataSource = new BasicDataSource();
        // please change host and port maybe like "jdbc:mysql://127.0.0.1:3306/test"
        dataSource.setUrl("jdbc:mysql://127.0.0.1:3306/test");
        dataSource.setUsername("root");
        dataSource.setPassword("qwer1234");
        // mysql schema, the sub schema for rootSchema, "test" is a schema in mysql
        Schema schema = JdbcSchema.create(rootSchema, "test", dataSource, null, "test");
        rootSchema.add("test", schema);
        // run sql query
        Statement statement = calciteConnection.createStatement();
        ResultSet resultSet = statement.executeQuery("select * from test.students");
        while (resultSet.next()) {
            System.out.println(resultSet.getObject(1) + "__" + resultSet.getObject(2));
        }
        statement.close();
        connection.close();
    }
}