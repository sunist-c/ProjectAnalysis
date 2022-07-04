package cn.sunist.server.bean;

import java.util.Date;
import java.util.Objects;

public class Country {
    private String countryName;
    private String refreshTime;
    private String confirm;
    private String death;
    private String recovered;

    public Country() {
    }

    public Country(String countryName, String refreshTime, String confirm, String death, String recovered) {
        this.countryName = countryName;
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
        Country country = (Country) o;
        return Objects.equals(countryName, country.countryName) && Objects.equals(refreshTime, country.refreshTime) && Objects.equals(confirm, country.confirm) && Objects.equals(death, country.death) && Objects.equals(recovered, country.recovered);
    }

    @Override
    public int hashCode() {
        return Objects.hash(countryName, refreshTime, confirm, death, recovered);
    }

    @Override
    public String toString() {
        return "Country{" +
                "countryName='" + countryName + '\'' +
                ", refreshTime='" + refreshTime + '\'' +
                ", confirm='" + confirm + '\'' +
                ", death='" + death + '\'' +
                ", recovered='" + recovered + '\'' +
                '}';
    }
}
