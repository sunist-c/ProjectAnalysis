package cn.sunist.server.dao;


import cn.sunist.server.bean.Country;
import cn.sunist.server.bean.Data;
import cn.sunist.server.bean.Province;
import cn.sunist.server.util.DateFormat;

import java.sql.*;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.List;

public class HiveJDBC {
    private static String url = "jdbc:hive2://10.42.1.60:10000/ProjectAnalysis";
    private static Connection con = null;

    static{
        try {
            Class.forName("org.apache.hive.jdbc.HiveDriver");
            con = DriverManager.getConnection(url,"root","123456");
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    private static void close(PreparedStatement ps,ResultSet resultSet) throws SQLException {
        try {
            if(resultSet != null){
                resultSet.close();
            }
        }catch (Exception e){}
        try {
            if(ps != null){
                ps.close();
            }
        }catch (Exception e){}
    }

    public static List<Object> calculateProvince(Date date, Integer time) throws SQLException {
        PreparedStatement ps = null;
        if(time == -1){
            String sql = "select location_country,location_province,sum(daily_confirm),sum(daily_death),sum(daily_recovered) from data where refresh_time <= ? group by location_country,location_province";
            ps = con.prepareStatement(sql);
            ps.setString(1, DateFormat.transform(date));
        }else{
            String sql = "select location_country,location_province,sum(daily_confirm),sum(daily_death),sum(daily_recovered) from data where ? < refresh_time and refresh_time <= ? group by location_country,location_province";
            ps = getPreparedStatement(date, time, sql);
        }
        return getResultSetProvince(date, ps);
    }

    private static PreparedStatement getPreparedStatement(Date date, Integer time, String sql) throws SQLException {
        PreparedStatement ps;
        ps = con.prepareStatement(sql);
        Calendar calendar = Calendar.getInstance();
        calendar.setTime(date);
        calendar.add(Calendar.DATE,-1 * time);
        ps.setString(1, DateFormat.transform(calendar.getTime()));
        ps.setString(2, DateFormat.transform(date));
        return ps;
    }

    public static List<Object> calculateCountry(Date date, Integer time) throws SQLException {
        PreparedStatement ps = null;
        if(time == -1){
            String sql = "select location_country,sum(daily_confirm),sum(daily_death),sum(daily_recovered) from data where refresh_time <= ? group by location_country";
            ps = con.prepareStatement(sql);
            ps.setString(1, DateFormat.transform(date));
        }else{
            String sql = "select location_country,sum(daily_confirm),sum(daily_death),sum(daily_recovered) from data where ? < refresh_time and refresh_time <= ? group by location_country";
            ps = getPreparedStatement(date, time, sql);
        }
        return getResultSetCountry(date, ps);
    }

    private static List<Object> getResultSetCountry(Date date, PreparedStatement ps) throws SQLException {
        ResultSet resultSet = ps.executeQuery();
        ArrayList<Object> list = new ArrayList<>();
        while (resultSet.next()){
            Country country = new Country(resultSet.getString(1),DateFormat.transform(date),resultSet.getString(2),resultSet.getString(3),resultSet.getString(4));
            list.add(country);
        }
        close(ps,resultSet);
        return list;
    }

    private static List<Object> getResultSetProvince(Date date, PreparedStatement ps) throws SQLException {
        ResultSet resultSet = ps.executeQuery();
        ArrayList<Object> list = new ArrayList<>();
        while (resultSet.next()){
            Province province = new Province(resultSet.getString(1), resultSet.getString(2), DateFormat.transform(date), resultSet.getString(3), resultSet.getString(4), resultSet.getString(5));
            list.add(province);
        }
        close(ps,resultSet);
        return list;
    }

    public static Boolean addData(List<Data> data) throws Exception {
        String dst = "/hive/temp.csv";
        HadoopJDBC.createFile(dst,data);
        String sql = "load data inpath "+ "'" + dst + "'"+ "into table data";
        PreparedStatement statement = con.prepareStatement(sql);
        boolean execute = statement.execute();
        HadoopJDBC.deleteFile(dst);
        return execute;
    }
}
