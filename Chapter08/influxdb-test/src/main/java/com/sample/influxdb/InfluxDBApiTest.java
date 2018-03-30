package com.sample.influxdb;

import java.util.concurrent.TimeUnit;

import org.influxdb.InfluxDB;
import org.influxdb.InfluxDBFactory;
import org.influxdb.dto.Point;
import org.influxdb.dto.Query;

public class InfluxDBApiTest {

	public static void main(String[] args) throws Exception {
		InfluxDB influxDB = InfluxDBFactory.connect("http://192.168.1.173:8086", "root", "root");
		String dbName = "testdb";
		influxDB.createDatabase(dbName);
		influxDB.setDatabase(dbName);
		Point point1 = Point.measurement("apachelog_java").time(1511361120000000000L, TimeUnit.NANOSECONDS)
				.addField("http", "GET /cgi-bin/try/ HTTP/1.0").addField("status", 200).addField("duration", 3395)
				.addField("ip", "192.168.2.20").build();
		Point point2 = Point.measurement("apachelog_java").time(1511361180000000000L, TimeUnit.NANOSECONDS)
				.addField("http", "GET / HTTP/1.0").addField("status", 200).addField("duration", 2216)
				.addField("ip", "127.0.0.1").build();
		influxDB.write(point1);
		influxDB.write(point2);
		Query query = new Query("SELECT * FROM apachelog_java", dbName);
		influxDB.query(query, 20, queryResult -> System.out.println(queryResult));

	}
}
