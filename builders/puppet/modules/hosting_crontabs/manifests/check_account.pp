class hosting_crontabs::check_account {

    cron { check_account:
        command => 'for acc in /var/cagefs/*/* ; do   test -e  ${acc%.lock}/etc/resolv.conf || echo  ${acc##*/} ; done',
        user    => root,
        minute  => "*/5",
    }

}
