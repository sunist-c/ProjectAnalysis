package cm.sunist.hadoop.bean;

import java.util.Date;
import java.util.Objects;

public class Country {
    private String countryName;
    private Date refreshTime;
    private Integer confirm;
    private Integer death;
    private Integer recovered;

    public Country() {
    }

    public Country(String countryName, Date refreshTime, Integer confirm, Integer death, Integer recovered) {
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

    public Date getRefreshTime() {
        return refreshTime;
    }

    public void setRefreshTime(Date refreshTime) {
        this.refreshTime = refreshTime;
    }

    public Integer getConfirm() {
        return confirm;
    }

    public void setConfirm(Integer confirm) {
        this.confirm = confirm;
    }

    public Integer getDeath() {
        return death;
    }

    public void setDeath(Integer death) {
        this.death = death;
    }

    public Integer getRecovered() {
        return recovered;
    }

    public void setRecovered(Integer recovered) {
        this.recovered = recovered;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Country that = (Country) o;
        return Objects.equals(countryName, that.countryName) && Objects.equals(refreshTime, that.refreshTime) && Objects.equals(confirm, that.confirm) && Objects.equals(death, that.death) && Objects.equals(recovered, that.recovered);
    }

    @Override
    public int hashCode() {
        return Objects.hash(countryName, refreshTime, confirm, death, recovered);
    }

    @Override
    public String toString() {
        return "Country{" +
                "countryName='" + countryName + '\'' +
                ", refreshTime=" + refreshTime +
                ", confirm=" + confirm +
                ", death=" + death +
                ", recovered=" + recovered +
                '}';
    }
}
