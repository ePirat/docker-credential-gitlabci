/* Copyright (c) 2022 Marvin Scholz <epirat07 at gmail dot com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package credhelper

import "testing"

func TestParseRegistryURL(t *testing.T) {
	urls := map[string]string{
		"gitlab.com":                  "gitlab.com",
		"http://gitlab.com":           "gitlab.com",
		"https://gitlab.com/":         "gitlab.com",
		"https://gitlab.com/foo/bar/": "gitlab.com",
		"registry.foo.bar:8080":       "registry.foo.bar",
		"[2001:DB8::]":                "2001:DB8::",
		"[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:8080": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	}
	for urlString, expectedHost := range urls {
		url, err := parseRegistryURL(urlString)
		if err != nil {
			t.Errorf("Got error while parsing URL: '%s'", url)
		}
		if url.Hostname() != expectedHost {
			t.Errorf("Parsed host '%s' did not match expected host: '%s'",
				url.Hostname(), expectedHost)
		}
	}
}

func checkMatch(t *testing.T, envURL string, servURL string, match bool) {
	t.Setenv("CI_REGISTRY", envURL)
	err := matchRegistryURL(servURL)
	if err != nil && match {
		t.Errorf("Error matching URLs '%s' and '%s' even though they should match.",
			envURL, servURL)
	} else if err == nil && !match {
		t.Errorf("Unexpected match of '%s' and '%s' even though they should NOT match.",
			envURL, servURL)
	}
}

func TestMatchRegistryURL(t *testing.T) {
	// Matches
	checkMatch(t, "registry.example.org", "registry.example.org", true)
	checkMatch(t, "https://registry.example.org", "registry.example.org", true)
	checkMatch(t, "registry.example.org", "https://registry.example.org", true)

	// Non-matches
	checkMatch(t, "registry.example.org", "registry.example.com", false)
	checkMatch(t, "https://registry.example.org", "registry.example.com", false)
	checkMatch(t, "registry.example.org", "https://registry.example.com", false)

	// Non-match of missmatching protocols
	checkMatch(t, "http://registry.example.org", "https://registry.example.org", false)

	// These must NOT match as the hosts are not valid
	checkMatch(t, "https://", "https://", false)
	checkMatch(t, "https://", "", false)
	checkMatch(t, "", "", false)
}
