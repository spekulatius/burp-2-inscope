<?php

try {
    // Prep the data
    $filename = $argv[1];
    // $mode = $argv[2];
    $url = 'https://about.gitlab.com';
    $parsedUrl = parse_url($url);

    // Get the JSON
    $json = json_decode(file_get_contents($filename), true);
    $includes = $json['target']['scope']['include'];
    $excludes = $json['target']['scope']['exclude'];

    // Start checking...
    $isIncluded = false;
    $isExcluded = false;

    foreach ($includes as $include) {
        // Check protocol
        if (isset($parsedUrl['scheme']) && $parsedUrl['scheme'] !== $include['protocol']) {
            continue;
        }

        // Check the port (checked only if set explicity).
        if (isset($parsedUrl['port']) && '^' . $parsedUrl['port'] . '$' !== $include['port']) {
            continue;
        }

        // Finally check if host matches.
        if (preg_match('/' . $include['host'] . '/', $parsedUrl['host'])) {
            $isIncluded = true;
        }
    }

    foreach ($excludes as $exclude) {
        // Check protocol
        if (isset($parsedUrl['scheme']) && $parsedUrl['scheme'] !== $exclude['protocol']) {
            continue;
        }

        // Check the port (checked only if set explicity).
        if (isset($parsedUrl['port']) && '^' . $parsedUrl['port'] . '$' !== $exclude['port']) {
            continue;
        }

        // Finally check if host matches.
        if (preg_match('/' . $exclude['host'] . '/', $parsedUrl['host'])) {
            $isExcluded = true;
        }
    }

    // Only consider the host further if it has been included but hasn't been excluded.
    if ($isIncluded && !$isExcluded) {
        echo $url;
    }
} catch (\Exception $e) {
    die($e);
}