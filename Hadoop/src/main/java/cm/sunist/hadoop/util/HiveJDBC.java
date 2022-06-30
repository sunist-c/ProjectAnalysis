package cm.sunist.hadoop.util;


import cm.sunist.hadoop.bean.Country;
import cm.sunist.hadoop.bean.Data;

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
            String sql = "select location_country,location_province,sum(daily_confirm),sum(daily_death),sum(daily_recovered) from data where ? < rerefresh_time and refresh_time <= ? group by location_country,location_province";
            ps = getPreparedStatement(date, time, sql);
        }
        return getResultSet(date, ps);
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
            String sql = "select location_country,sum(daily_confirm),sum(daily_death),sum(daily_recovered) from data where ? < rerefresh_time and refresh_time <= ? group by location_country";
            ps = getPreparedStatement(date, time, sql);
        }
        return getResultSet(date, ps);
    }

    private static List<Object> getResultSet(Date date, PreparedStatement ps) throws SQLException {
        ResultSet resultSet = ps.executeQuery();
        ArrayList<Object> list = new ArrayList<>();
        while (resultSet.next()){
            Country country = new Country(resultSet.getString(1),date,Integer.valueOf(resultSet.getString(2)),Integer.valueOf(resultSet.getString(3)),Integer.valueOf(resultSet.getString(4)));
            list.add(country);
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
