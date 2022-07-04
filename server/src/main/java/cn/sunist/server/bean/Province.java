package cn.sunist.server.bean;

import java.util.Date;
import java.util.Objects;

public class Province {
    private String countryName;
    private String provinceName;
    private String refreshTime;
    private String confirm;
    private String death;
    private String recovered;

    public Province() {
    }

    public Province(String countryName, String provinceName, String refreshTime, String confirm, String death, String recovered) {
        this.countryName = countryName;
        this.provinceName = provinceName;
        this.refreshTime = refreshTime;
        this.confirm = confirm;
        this.death = death;
        this.recovered = recovered;
    }

    public String getCountryName() {
        return countryName;
    }

    public void setCountryName(String countryName) {
        this.countryName = countryName;
    }

    public String getProvinceName() {
        return provinceName;
    }

    public void setProvinceName(String provinceName) {
        this.provinceName = provinceName;
    }

    public String getRefreshTime() {
        return refreshTime;
    }

    public void setRefreshTime(String refreshTime) {
        this.refreshTime = refreshTime;
    }

    public String getConfirm() {
        return confirm;
    }

    public void setConfirm(String confirm) {
        this.confirm = confirm;
    }

    public String getDeath() {
        return death;
    }

    public void setDeath(String death) {
        this.death = death;
    }

    public String getRecovered() {
        return recovered;
    }

    public void setRecovered(String recovered) {
        this.recovered = recovered;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Province province = (Province) o;
        return Objects.equals(countryName, province.countryName) && Objects.equals(provinceName, province.provinceName) && Objects.equals(refreshTime, province.refreshTime) && Objects.equals(confirm, province.confirm) && Objects.equals(death, province.death) && Objects.equals(recovered, province.recovered);
    }

    @Override
    public int hashCode() {
        return Objects.hash(countryName, provinceName, refreshTime, confirm, death, recovered);
    }

    @Override
    public String toString() {
        return "Province{" +
                "countryName='" + countryName + '\'' +
                ", provinceName='" + provinceName + '\'' +
                ", refreshTime='" + refreshTime + '\'' +
                ", confirm='" + confirm + '\'' +
                ", death='" + death + '\'' +
                ", recovered='" + recovered + '\'' +
                '}';
    }
}
