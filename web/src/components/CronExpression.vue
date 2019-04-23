<template>
  <div id="changeContab">
    <el-tabs type="border-card">
      <el-tab-pane>
        <span slot="label">Seconds</span>
        <div class="tabBody">
          <el-row>
            <el-radio v-model="second.cronEvery" label="1">Every second</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="second.cronEvery" label="2">Every
              <el-input-number size="small" v-model="second.incrementIncrement" :min="1" :max="60" :disabled="second.cronEvery !== '2'"></el-input-number>
              second(s) starting at second
              <el-input-number size="small" v-model="second.incrementStart" :min="0" :max="59" :disabled="second.cronEvery !== '2'"></el-input-number>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="second.cronEvery" label="3">Specific second (choose one or many)
              <el-select placeholder="Please select" size="small" multiple v-model="second.specificSpecific" :disabled="second.cronEvery !== '3'">
                <el-option v-for="val in 60" :key="val" :value="val-1">{{val-1}}</el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="second.cronEvery" label="4">Every second between second
              <el-input-number size="small" v-model="second.rangeStart" :min="1" :max="60" :disabled="second.cronEvery !== '4'"></el-input-number>
              and second
              <el-input-number size="small" v-model="second.rangeEnd" :min="0" :max="59" :disabled="second.cronEvery !== '4'"></el-input-number>
            </el-radio>
          </el-row>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">Minutes</span>
        <div class="tabBody">
          <el-row>
            <el-radio v-model="minute.cronEvery" label="1">Every minute</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="minute.cronEvery" label="2">Every
              <el-input-number size="small" v-model="minute.incrementIncrement" :min="1" :max="60" :disabled="minute.cronEvery !== '2'"></el-input-number>
              minute(s) starting at minute
              <el-input-number size="small" v-model="minute.incrementStart" :min="0" :max="59" :disabled="minute.cronEvery !== '2'"></el-input-number>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="minute.cronEvery" label="3">Specific minute (choose one or many)
              <el-select size="small" multiple v-model="minute.specificSpecific" :disabled="minute.cronEvery !== '3'">
                <el-option v-for="val in 60" :key="val" :value="val-1">{{val-1}}</el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="minute.cronEvery" label="4">Every minute between minute
              <el-input-number size="small" v-model="minute.rangeStart" :min="1" :max="60" :disabled="minute.cronEvery !== '4'"></el-input-number>
              and minute
              <el-input-number size="small" v-model="minute.rangeEnd" :min="0" :max="59" :disabled="minute.cronEvery !== '4'"></el-input-number>
            </el-radio>
          </el-row>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">Hours</span>
        <div class="tabBody">
          <el-row>
            <el-radio v-model="hour.cronEvery" label="1">Every hour</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="hour.cronEvery" label="2">Every
              <el-input-number size="small" v-model="hour.incrementIncrement" :min="0" :max="23" :disabled="hour.cronEvery !== '2'"></el-input-number>
              hour(s) starting at hour
              <el-input-number size="small" v-model="hour.incrementStart" :min="0" :max="23" :disabled="hour.cronEvery !== '2'"></el-input-number>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="hour.cronEvery" label="3">Specific hour (choose one or many)
              <el-select size="small" multiple v-model="hour.specificSpecific" :disabled="hour.cronEvery !== '3'">
                <el-option v-for="val in 24" :key="val" :value="val-1">{{val-1}}</el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="hour.cronEvery" label="4">Every hour between hour
              <el-input-number size="small" v-model="hour.rangeStart" :min="0" :max="23" :disabled="hour.cronEvery !== '4'"></el-input-number>
              and hour
              <el-input-number size="small" v-model="hour.rangeEnd" :min="0" :max="23" :disabled="hour.cronEvery !== '4'"></el-input-number>
            </el-radio>
          </el-row>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">Day</span>
        <div class="tabBody">
          <el-row>
            <el-radio v-model="day.cronEvery" label="1">Every day</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="2">Every
              <el-input-number size="small" v-model="week.incrementIncrement" :min="1" :max="7" :disabled="day.cronEvery !== '2'"></el-input-number>
              day(s) starting on
              <el-select size="small" v-model="week.incrementStart" :disabled="day.cronEvery !== '2'">
                <el-option v-for="week in options.week" :key="week.value" :label="week.label" :value="week.value"></el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="3">Every
              <el-input-number size="small" v-model="day.incrementIncrement" :min="1" :max="31" :disabled="day.cronEvery !== '3'"></el-input-number>
              day(s) starting at the','of the month
              <el-input-number size="small" v-model="day.incrementStart" :min="1" :max="31" :disabled="day.cronEvery !== '3'"></el-input-number>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="day.cronEvery" label="4">Specific day of week (choose one or many)
              <el-select size="small" multiple v-model="week.specificSpecific" :disabled="day.cronEvery !== '4'">
                <el-option v-for="week in options.week" :key="week.value" :label="week.label" :value="week.short"></el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="day.cronEvery" label="5">Specific day of month (choose one or many)
              <el-select size="small" multiple v-model="day.specificSpecific" :disabled="day.cronEvery !== '5'">
                <el-option v-for="val in 31" :key="val" :value="val">{{val}}</el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="6">On the last day of the month</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="7">On the last weekday of the month</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="8">On the last
              <el-select size="small" v-model="day.cronLastSpecificDomDay" :disabled="day.cronEvery !== '8'">
                <el-option v-for="week in options.week" :key="week.value" :label="week.label" :value="week.value"></el-option>
              </el-select>
              of the month
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="9">
              <el-input-number size="small" v-model="day.cronDaysBeforeEomMinus" :min="1" :max="31" :disabled="day.cronEvery !== '9'"></el-input-number>
              day(s) before the end of the month
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="10">Nearest weekday (Monday to Friday) to the
              <el-input-number size="small" v-model="day.cronDaysNearestWeekday" :min="1" :max="31" :disabled="day.cronEvery !== '10'"></el-input-number>
              of the month
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="day.cronEvery" label="11">On the
              <el-input-number size="small" v-model="week.cronNthDayNth" :min="1" :max="5" :disabled="day.cronEvery !== '11'"></el-input-number>
              <el-select size="small" v-model="week.cronNthDayDay" :disabled="day.cronEvery !== '11'">
                <el-option v-for="week in options.week" :key="week.value" :label="week.label" :value="week.value"></el-option>
              </el-select>
              of the month
            </el-radio>
          </el-row>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">Month</span>
        <div class="tabBody">
          <el-row>
            <el-radio v-model="month.cronEvery" label="1">Every month</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="month.cronEvery" label="2">Every
              <el-input-number size="small" v-model="month.incrementIncrement" :min="0" :max="12" :disabled="month.cronEvery !== '2'"></el-input-number>
              month(s) starting in
              <el-input-number size="small" v-model="month.incrementStart" :min="0" :max="12" :disabled="month.cronEvery !== '2'"></el-input-number>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="month.cronEvery" label="3">Specific month (choose one or many)
              <el-select size="small" multiple v-model="month.specificSpecific" :disabled="month.cronEvery !== '3'">
                <el-option v-for="val in 12" :key="val" :label="val" :value="val"></el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="month.cronEvery" label="4">Every month between
              <el-input-number size="small" v-model="month.rangeStart" :min="1" :max="12" :disabled="month.cronEvery !== '4'"></el-input-number>
              and
              <el-input-number size="small" v-model="month.rangeEnd" :min="1" :max="12" :disabled="month.cronEvery !== '4'"></el-input-number>
            </el-radio>
          </el-row>
        </div>
      </el-tab-pane>
      <el-tab-pane>
        <span slot="label">Year</span>
        <div class="tabBody">
          <el-row>
            <el-radio v-model="year.cronEvery" label="1">Any year</el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="year.cronEvery" label="2">Every
              <el-input-number size="small" v-model="year.incrementIncrement" :min="1" :max="99" :disabled="year.cronEvery !== '2'"></el-input-number>
              year(s) starting in
              <el-input-number size="small" v-model="year.incrementStart" :min="2018" :max="2118" :disabled="year.cronEvery !== '2'"></el-input-number>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio class="long" v-model="year.cronEvery" label="3">Specific year (choose one or many)
              <el-select size="small" filterable multiple v-model="year.specificSpecific" :disabled="year.cronEvery !== '3'">
                <el-option v-for="val in 100" :key="val" :label="2017+val" :value="2017+val"></el-option>
              </el-select>
            </el-radio>
          </el-row>
          <el-row>
            <el-radio v-model="year.cronEvery" label="4">Every year between
              <el-input-number size="small" v-model="year.rangeStart" :min="time.getFullYear()" :max="2118" :disabled="year.cronEvery !== '4'"></el-input-number>
              and
              <el-input-number size="small" v-model="year.rangeEnd" :min="time.getFullYear()" :max="2118" :disabled="year.cronEvery !== '4'"></el-input-number>
            </el-radio>
          </el-row>
        </div>
      </el-tab-pane>
    </el-tabs>
    <div class="bottom">
      <el-button @click="change">Cancel</el-button>
      <el-button type="primary" @click="confirm">Confirm</el-button>
    </div>
  </div>
</template>
<script>
  export default {
    name: 'vueCron',
    props: ['data'],
    data() {
      return {
        time: new Date(),
        options: {
          week: [
            {
              value: 7,
              short: "SUN",
              label: "Sunday"
            },{
              value: 1,
              short: "MON",
              label: "Monday"
            },{
              value: 2,
              short: "TUE",
              label: "Tuesday"
            },{
              value: 3,
              short: "WED",
              label: "Wednesday"
            },{
              value: 4,
              short: "THU",
              label: "Thursday"
            },{
              value: 5,
              short: "FRI",
              label: "Friday"
            },{
              value: 6,
              short: "SAT",
              label: "Saturday"
            },
          ]
        },
        second: {
          cronEvery: '1',
          incrementStart: '3',
          incrementIncrement: '5',
          rangeStart: '',
          rangeEnd: '',
          specificSpecific: [],
        },
        minute: {
          cronEvery: '1',
          incrementStart: '3',
          incrementIncrement: '5',
          rangeStart: '',
          rangeEnd: '',
          specificSpecific: [],
        },
        hour: {
          cronEvery: '1',
          incrementStart: '3',
          incrementIncrement: '5',
          rangeStart: '',
          rangeEnd: '',
          specificSpecific: [],
        },
        day: {
          cronEvery: '1',
          incrementStart: '1',
          incrementIncrement: '1',
          rangeStart: '',
          rangeEnd: '',
          specificSpecific: [],
          cronLastSpecificDomDay: 1,
          cronDaysBeforeEomMinus: '',
          cronDaysNearestWeekday: '',
        },
        week: {
          cronEvery: '1',
          incrementStart: '1',
          incrementIncrement: '1',
          specificSpecific: [],
          cronNthDayDay: 1,
          cronNthDayNth: '1',
        },
        month: {
          cronEvery: '1',
          incrementStart: '3',
          incrementIncrement: '5',
          rangeStart: '',
          rangeEnd: '',
          specificSpecific: [],
        },
        year: {
          cronEvery: '1',
          incrementStart: '2017',
          incrementIncrement: '1',
          rangeStart: '',
          rangeEnd: '',
          specificSpecific: [],
        },
        output: {
          second: '',
          minute: '',
          hour: '',
          day: '',
          month: '',
          Week: '',
          year: '',
        }
      }
    },
    watch: {
      data() {
        this.rest(this.$data);
      }
    },
    computed: {
      secondsText() {
        let seconds = '';
        let cronEvery = this.second.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
            seconds = '*';
            break;
          case '2':
            seconds = this.second.incrementStart + '/' + this.second.incrementIncrement;
            break;
          case '3':
            this.second.specificSpecific.map(val => {
              seconds += val + ','
            });
            seconds = seconds.slice(0, -1);
            break;
          case '4':
            seconds = this.second.rangeStart + '-' + this.second.rangeEnd;
            break;
        }
        return seconds;
      },
      minutesText() {
        let minutes = '';
        let cronEvery = this.minute.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
            minutes = '*';
            break;
          case '2':
            minutes = this.minute.incrementStart + '/' + this.minute.incrementIncrement;
            break;
          case '3':
            this.minute.specificSpecific.map(val => {
              minutes += val + ','
            });
            minutes = minutes.slice(0, -1);
            break;
          case '4':
            minutes = this.minute.rangeStart + '-' + this.minute.rangeEnd;
            break;
        }
        return minutes;
      },
      hoursText() {
        let hours = '';
        let cronEvery = this.hour.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
            hours = '*';
            break;
          case '2':
            hours = this.hour.incrementStart + '/' + this.hour.incrementIncrement;
            break;
          case '3':
            this.hour.specificSpecific.map(val => {
              hours += val + ','
            });
            hours = hours.slice(0, -1);
            break;
          case '4':
            hours = this.hour.rangeStart + '-' + this.hour.rangeEnd;
            break;
        }
        return hours;
      },
      daysText() {
        let days = '';
        let cronEvery = this.day.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
            break;
          case '2':
          case '4':
          case '11':
            days = '?';
            break;
          case '3':
            days = this.day.incrementStart + '/' + this.day.incrementIncrement;
            break;
          case '5':
            this.day.specificSpecific.map(val => {
              days += val + ','
            });
            days = days.slice(0, -1);
            break;
          case '6':
            days = "L";
            break;
          case '7':
            days = "LW";
            break;
          case '8':
            days = this.day.cronLastSpecificDomDay + 'L';
            break;
          case '9':
            days = 'L-' + this.day.cronDaysBeforeEomMinus;
            break;
          case '10':
            days = this.day.cronDaysNearestWeekday + "W";
            break
        }
        return days;
      },
      weeksText() {
        let weeks = '';
        let cronEvery = this.day.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
          case '3':
          case '5':
            weeks = '?';
            break;
          case '2':
            weeks = this.week.incrementStart + '/' + this.week.incrementIncrement;
            break;
          case '4':
            this.week.specificSpecific.map(val => {
              weeks += val + ','
            });
            weeks = weeks.slice(0, -1);
            break;
          case '6':
          case '7':
          case '8':
          case '9':
          case '10':
            weeks = "?";
            break;
          case '11':
            weeks = this.week.cronNthDayDay + "#" + this.week.cronNthDayNth;
            break;
        }
        return weeks;
      },
      monthsText() {
        let months = '';
        let cronEvery = this.month.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
            months = '*';
            break;
          case '2':
            months = this.month.incrementStart + '/' + this.month.incrementIncrement;
            break;
          case '3':
            this.month.specificSpecific.map(val => {
              months += val + ','
            });
            months = months.slice(0, -1);
            break;
          case '4':
            months = this.month.rangeStart + '-' + this.month.rangeEnd;
            break;
        }
        return months;
      },
      yearsText() {
        let years = '';
        let cronEvery = this.year.cronEvery;
        switch (cronEvery.toString()) {
          case '1':
            years = '*';
            break;
          case '2':
            years = this.year.incrementStart + '/' + this.year.incrementIncrement;
            break;
          case '3':
            this.year.specificSpecific.map(val => {
              years += val + ','
            });
            years = years.slice(0, -1);
            break;
          case '4':
            years = this.year.rangeStart + '-' + this.year.rangeEnd;
            break;
        }
        return years;
      },
      cron() {
        return `${this.secondsText || '*'} ${this.minutesText || '*'} ${this.hoursText || '*'} ${this.daysText || '*'} ${this.monthsText || '*'} ${this.weeksText || '?'} ${this.yearsText || '*'}`
      },
    },
    methods: {
      change() {
        this.$emit('change', this.cron);
        this.$emit('close');
      },
      confirm() {
        this.change();
        this.$emit('close');
      },
      rest(data) {
        for (let i in data) {
          if (data[i] instanceof Object) {
            this.rest(data[i])
          } else {
            switch (typeof data[i]) {
              case 'object':
                data[i] = [];
                break;
              case 'string':
                data[i] = '';
                break;
            }
          }
        }
      }
    },
    mounted() {
    }
  }
</script>

<style lang="scss" scoped>
  #changeContab {
    .el-tabs {
      box-shadow: none;
    }

    .tabBody {
      .el-row {
        margin: 10px 0;
        .el-input-number {
          width: 110px;
        }

      }
    }

    .bottom {
      width: 100%;
      text-align: center;
      margin: 20px 0 8px 0;
      position: relative;

      .value {
        font-size: 18px;
        vertical-align: middle;
      }

    }
  }
</style>
