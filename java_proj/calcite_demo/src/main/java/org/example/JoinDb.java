package org.example;

import org.apache.calcite.adapter.jdbc.JdbcSchema;
import org.apache.calcite.jdbc.CalciteConnection;
import org.apache.calcite.schema.Schema;
import org.apache.calcite.schema.SchemaPlus;
import org.apache.commons.dbcp2.BasicDataSource;

import java.sql.*;
import java.util.Properties;

public class JoinDb {
    public static void main(String[] args) throws ClassNotFoundException, SQLException {
        System.out.println("Hello join db!");
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
        // mysql database 1
        BasicDataSource dataSource = new BasicDataSource();
        // please change host and port maybe like "jdbc:mysql://127.0.0.1:3306/test"
        dataSource.setUrl("jdbc:mysql://127.0.0.1:3306/test");
        dataSource.setUsername("root");
        dataSource.setPassword("qwer1234");
        // mysql schema, the sub schema for rootSchema, "test" is a schema in mysql
        Schema schema = JdbcSchema.create(rootSchema, "test", dataSource, null, "test");
        rootSchema.add("test", schema);

        // mysql database 2
        BasicDataSource dataSource2 = new BasicDataSource();
        // please change host and port maybe like "jdbc:mysql://127.0.0.1:3306/test"
        dataSource2.setUrl("jdbc:mysql://127.0.0.1:3306/test2");
        dataSource2.setUsername("root");
        dataSource2.setPassword("qwer1234");
        // mysql schema, the sub schema for rootSchema, "test" is a schema in mysql
        Schema schema2 = JdbcSchema.create(rootSchema, "test2", dataSource2, null, "test2");
        rootSchema.add("test2", schema2);





        // run sql query
        Statement statement = calciteConnection.createStatement();
        ResultSet resultSet = statement.executeQuery("select * from test.students " +
                "left join test2.gifts on test.students.id = test2.gifts.id");
        // ResultSet resultSet = statement.executeQuery("select * from test2.gifts");
        while (resultSet.next()) {
            System.out.println(resultSet.getObject(1) +
                    "__" + resultSet.getObject(2)+
                    "__" + resultSet.getObject(3)+
                    "__" + resultSet.getObject(4));
        }
        statement.close();
        connection.close();
    }
}